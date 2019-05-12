package fe26

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/AndresBott/Fe26/pkg/f"
	"path/filepath"
	"github.com/AndresBott/Fe26/internal/fileManager"
	"strings"
	"errors"
	"encoding/json"
	"net/url"
	"github.com/dustin/go-humanize"
	"github.com/Masterminds/sprig"
	"html/template"
	"github.com/gobuffalo/packr/v2"
)

type Breadcrumb struct{
	Name string
	Url string
}

type listFileData  struct {
	Files []fileManager.File
	Dirs []fileManager.File
	RequestUrl string
	RequestBasePath string
	RequestParentPath string
	BreadCrumbs []Breadcrumb
}

// take a relative directory path and separate breadcrumb types
// return a sequencial array of breadcrumbs
func breadcrumbs(directoryQueryPath string) []Breadcrumb  {

	pathParts := strings.Split(directoryQueryPath,"/")
	var newPath []Breadcrumb
	breadCrumbsUrl := ""
	for _, item := range pathParts {
		if item != ""{
			breadCrumbsUrl = breadCrumbsUrl + "/"+item
			b := Breadcrumb{
				Name:item,
				Url:breadCrumbsUrl,
			}
			newPath = append(newPath, b)
		}
	}
	return newPath
}

// list the files of a dir based on the http query string
func getFilesData(directoryQueryPath string) (listFileData,error) {

	data := listFileData{}

	if directoryQueryPath == ""{
		directoryQueryPath = "/"
	}

	sanitizedPath,err := f.Sanitize(directoryQueryPath)
	if err != nil{
		data.RequestUrl="/"
		return data,err
	}

	if directoryQueryPath != sanitizedPath{
		data.RequestUrl=sanitizedPath
		return data,errors.New("Request malformed: "+directoryQueryPath)
	}

	data.RequestBasePath = ""
	if directoryQueryPath != "/"{
		data.RequestBasePath = directoryQueryPath
	}
	data.RequestUrl = directoryQueryPath
	data.RequestParentPath = filepath.Dir(directoryQueryPath)

	d := fileManager.NewFileManager(directoryQueryPath,Config.docRoot)
	data.Files = d.Files
	data.Dirs = d.Dirs

	data.BreadCrumbs = breadcrumbs(directoryQueryPath)

	return data,nil

}

func ListFilesHTML(w http.ResponseWriter, r *http.Request)  {
	d := r.URL.Query().Get("d")
	data,err := getFilesData(d)

	if err != nil{
		log.Warn("Wrong query parameter, redirecting to: "+data.RequestUrl)
		http.Redirect(w, r, Config.FeBase+"?d="+data.RequestUrl, 302)
		return
	}

	log.Info("LIST: "+data.RequestUrl)
	templates := packr.New("Templates", "../../web/templates")

	s, _ := templates.FindString("view.html")
	s2 ,_ := templates.FindString("head.html")
	s3 ,_ := templates.FindString("fileupload.html")

	tmpl,err := template.New("main").Funcs(
		// todo externalize funcMAp
		template.FuncMap{
			"urlEncode": func(uri string) string {
				//https://play.golang.org/p/pOfrn-Wsq5
				s := url.URL{Path: uri}
				return s.String()
				//return  url.QueryEscape(uri)
			},
			"humanizeBytes": func (b int64) string{
				return humanize.Bytes(uint64(b))
			},
		}).Funcs(sprig.FuncMap()).Parse(s+s2+s3)

	if err != nil{
		log.Warn(err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Error(err)
	}
}

func ListFilesJson(w http.ResponseWriter, r *http.Request)  {
	d := r.URL.Query().Get("d")
	data,err := getFilesData(d)

	// TODO: maybe json should not redirect
	if err != nil{
		log.Warn("Wrong query parameter, redirecting to: "+data.RequestUrl)
		http.Redirect(w, r, Config.FeBase+"?d="+data.RequestUrl, 302)
		return
	}

	log.Info("JSON LIST: "+data.RequestUrl)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, ServerError, http.StatusInternalServerError)
		log.Error("Unable to serve: "+data.RequestUrl+ " due to "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}