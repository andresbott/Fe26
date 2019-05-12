package fileManager

import (
	"os"
	"time"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"io/ioutil"
	"errors"
	f "github.com/AndresBott/f"
)

type File struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
	FileType string
}

type FileManager struct {
	Files []File
	Dirs []File
	Root string
}

func NewFileManager(url string, documentRoot string) FileManager {

	absDir, err := getAbsPath(url,documentRoot)

	if err != nil{
		log.Error(err)
	}
	d:= FileManager{Root:absDir}
	d.ReadDir()
	return d
}

func (d *FileManager) ReadDir() {
	files, err := ioutil.ReadDir(d.Root)
	if err != nil {
		log.Error(err)
	}

	for _, file := range files {

		nf := File{
			Name:    file.Name(),
			Size:    file.Size(),
			Mode:    file.Mode(),
			ModTime: file.ModTime(),
			IsDir:   file.IsDir(),
		}

		if file.IsDir(){
			nf.FileType = "folder"
			d.Dirs = append(d.Dirs, nf)
		}else{
			ext := filepath.Ext(nf.Name)
			if ext != "" {
				ext = ext[1:]
			}else {
				ext = ""
			}
			ft := f.FileTypeFromExtension(ext)
			nf.FileType = ft.Type

			d.Files = append(d.Files,nf)

		}
	}
}


func getAbsPath(url string,documentRoot string) (string,error)  {

	rel,_ := filepath.Rel( documentRoot, documentRoot+url)

	if strings.HasPrefix(rel, "..") {
		log.Warn("requested path tries to exit the document root, correcting")
		url = "/"
	}

	abs,_ := filepath.Abs(documentRoot+url)
	file,err := os.Stat(abs)

	if err != nil {
		return "",errors.New("Path does not exist: "+abs)
	}

	if file.Mode().IsDir() {
		return abs,nil
	}else {
		return "",errors.New("Not a directory: "+abs)
	}
}


