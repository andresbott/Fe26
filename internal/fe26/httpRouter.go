package fe26

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	log "github.com/sirupsen/logrus"
	"github.com/gobuffalo/packr/v2"
	"os"
	"strings"
	"path"
)


// localRedirect gives a Moved Permanently response.
// It does not convert relative paths to absolute paths like Redirect does.
func localRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	log.Info("REDIRECT: from: "+r.URL.Path+"to"+newPath)
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(StatusMovedPermanently)
}

func redirectToView(w http.ResponseWriter, r *http.Request) {
	localRedirect(w,r,Config.FeBase)
}

func isJsonRequest( r *http.Request) bool {
	return strings.HasSuffix(path.Base(r.URL.Path), "json")
}

func handlePost(w http.ResponseWriter, r *http.Request) {

	log.Debug("POST[Content-Type] = "+r.Header.Get("Content-Type"))

	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data"){
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			log.Error(err)
			http.Error(w, ServerError, http.StatusInternalServerError)
			return
		}

		action := r.FormValue("action")
		if action == "upload" {
			doFileUpload(w,r)
		}else {
			log.Error("Wrong action: "+action)
			http.Error(w, ServerBadRequest, http.StatusBadRequest)
			return
		}
	}

	// handle file delete
	//curl -d "action=delete&path=/" -X POST http://localhost:8080/fe26 -v
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded"){
		if err := r.ParseForm(); err != nil {
			log.Error(err)
			return
		}
		action := r.FormValue("action")
		if action == "delete-file" {
			doFileDelete(w,r)
		}else if action == "create-dir"{
			doCrateDir(w,r)
		}else {
			log.Error("Wrong action: "+action)
			http.Error(w, ServerBadRequest, http.StatusBadRequest)
			return
		}
	}
}



func fe26Router()  {

	r := mux.NewRouter()

	// Base path: i.e GET fe26 & fe26.json
	r.HandleFunc("/"+Config.FeBase, ListFilesHTML).Methods("GET")
	r.HandleFunc("/"+Config.FeBase+".json", ListFilesJson).Methods("GET")

	// Handle static files like css and js files
	staticFiles := packr.New("static", "../../web/static")
	r.PathPrefix("/static.file").Handler(http.StripPrefix("/static.file", http.FileServer( staticFiles ) )).Methods("GET")

	// if GET to / redirect to fe26
	r.HandleFunc("/", redirectToView ).Methods("GET")

	// Handle static files that exist
	r.Handle("/{filePath:.*}", http.FileServer( Fe26Dir(os.Getenv("FE26_ROOT") ) ) ).Methods("GET")

	// Handle Post requests
	//r.HandleFunc("/"+config.FeBase, handlePost).Methods("POST")
	r.HandleFunc("/"+Config.FeBase+".json", handlePost).Methods("POST")

	http.Handle("/", r)

	if err := http.ListenAndServe(":"+strconv.Itoa(Config.port), nil); err != nil {
		log.Fatal(err)
	}

}






