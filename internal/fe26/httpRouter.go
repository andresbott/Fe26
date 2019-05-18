package fe26

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	log "github.com/sirupsen/logrus"
	"github.com/gobuffalo/packr/v2"
	"os"
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
	//r.HandleFunc("/"+config.FeBase+".json", handlePost).Methods("POST")

	http.Handle("/", r)

	if err := http.ListenAndServe(":"+strconv.Itoa(Config.port), nil); err != nil {
		log.Fatal(err)
	}

}






