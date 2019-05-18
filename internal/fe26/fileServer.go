package fe26

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"errors"
	"path/filepath"
	"path"
	"os"
	"io"
	"encoding/json"
	"github.com/AndresBott/Fe26/pkg/f"
)

// Re implement http.Dir in order to not expose directories
type Fe26Dir string

// Open opens file
func (d Fe26Dir) Open(name string) (http.File, error) {

	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	fullName := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name)))
	f, err := os.Open(fullName)
	if err != nil {
		return nil, err
	}
	// check if is dir
	s, err := f.Stat()
	if s.IsDir() {
		log.Warn("FILE: tying to access folder: "+name+", 404 returned")
		return nil,os.ErrNotExist
	}
	log.Info("FILE:"+name)
	return f, nil
}



func doFileUpload(w http.ResponseWriter, r *http.Request){
	log.Debug("doFileUpload")

	path := r.FormValue("path")
	if path == ""{
		log.Error("UPLOAD: Destination Path not defined")
		http.Error(w, ServerBadRequest, http.StatusBadRequest)
		return
	}

	fhs := r.MultipartForm.File["files"]
	// TODO: add a max ammount of files allowed to upload
	for _, fileHandler := range fhs {
		log.Info("here2")
		file, err := fileHandler.Open()
		defer file.Close()

		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debug("try to upload: "+fileHandler.Filename+ "to path: "+Config.docRoot+path )

		// TODO: Not overwrite existing file
		f, err := os.OpenFile(Config.docRoot+path+"/"+fileHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Error("UPLOAD: "+err.Error())
			return
		}
		defer f.Close()
		io.Copy(f, file)

		log.Info("UPLOAD: "+Config.docRoot+path+"/"+fileHandler.Filename)
	}

	if isJsonRequest(r){
		var rdata struct{
			Success *bool `json:"success"`
		}
		t := new(bool)
		*t = true
		rdata.Success = t
		js, err := json.Marshal(rdata)
		if err != nil {
			http.Error(w, ServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}else {
		defer 	redirectToView(w,r)
	}
}


func doFileDelete(w http.ResponseWriter, r *http.Request){
	path := r.FormValue("filepath")
	if path == ""{
		log.Error("DELETE: File path is not defined")
		http.Error(w, ServerBadRequest, http.StatusBadRequest)
		return
	}

	err := os.Remove(Config.docRoot+path)
	if err != nil{
		log.Error("Unable to delete file: "+Config.docRoot+path)
		http.Error(w, ServerBadRequest, http.StatusBadRequest)
		return
	}else{
		log.Info("DELETE: "+path)
		log.Debug("DELETE - abs path: "+Config.docRoot+path)
	}
}

func doCrateDir(w http.ResponseWriter, r *http.Request){
	path := r.FormValue("path")
	if path == ""{
		log.Error("CREATE DIR: Destination Path not defined")
		http.Error(w, ServerBadRequest, http.StatusBadRequest)
		return
	}
	folderName := r.FormValue("foldername")

	pathToCreate := Config.docRoot+path+ "/" + folderName
	parentPath,err := os.Stat(filepath.Dir(pathToCreate))
	if err != nil{
		log.Error("Unable to read parent dir: "+filepath.Dir(pathToCreate))
		http.Error(w, ServerBadRequest, http.StatusBadRequest)
		return
	}

	log.Debug("Real dir path: "+pathToCreate+" with mode: "+parentPath.Mode().String())

	// TODO, add check if directory exists
	mkdirErr :=os.Mkdir(Config.docRoot+path+ "/" + folderName,parentPath.Mode())

	if mkdirErr != nil{
		log.Error("Error creating dir: "+Config.docRoot+path+ "/" + folderName+" msg: "+mkdirErr.Error())
	}else {
		fname,_ := f.Sanitize(path+ "/" + folderName)
		log.Info("CREATE DIR: "+ fname )
	}
}