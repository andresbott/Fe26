package fe26

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
)

func fe26Router()  {

	r := mux.NewRouter()

	r.HandleFunc("/"+Config.FeBase, ListFilesHTML).Methods("GET")
	r.HandleFunc("/"+Config.FeBase+".json", ListFilesJson).Methods("GET")


	//staticFiles := packr.New("static", "./UI/static")
	//r.PathPrefix("/static.file").Handler(http.StripPrefix("/static.file", http.FileServer( staticFiles ) )).Methods("GET")
	//
	////r.HandleFunc("/{filePath:.*}", serveFile)
	//fileServer := http.FileServer(FileSystem{http.Dir(os.Getenv("FE26_ROOT"))})
	//
	//r.HandleFunc("/", redirectToView ).Methods("GET")
	//
	//r.Handle("/{filePath:.*}",fileServer ).Methods("GET")
	//
	////http.Handle("/", fileServer )
	//r.HandleFunc("/"+config.FeBase, handlePost).Methods("POST")
	//r.HandleFunc("/"+config.FeBase+".json", handlePost).Methods("POST")
	//
	//
	http.Handle("/", r)

	if err := http.ListenAndServe(":"+strconv.Itoa(Config.port), nil); err != nil {
		log.Fatal(err)
	}

}