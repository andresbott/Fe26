package fe26

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type Handler struct {
	router *mux.Router
	root   string
}

func (fe26 *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fe26.router.ServeHTTP(w, r)
}

func New(l *zerolog.Logger, Root string) (*Handler, error) {
	r := mux.NewRouter()

	// add logging middleware
	//r.Use(func(handler http.Handler) http.Handler {
	//	return zero.LoggingMiddleware(handler, l)
	//})
	r.Path("/").Handler(HandleDirContent(""))
	return &Handler{
		router: r,
		root:   Root,
	}, nil
}

type FsEntry struct {
	Name string `json:"name"`
}
type Resp struct {
	Dirs  []FsEntry `json:"dirs"`
	Files []FsEntry `json:"files"`
}

func HandleDirContent(path string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: use proper abstraction
		osFiles, err := os.ReadDir("./")
		if err != nil {
			fmt.Printf("[ERROR]: %v", err)
		}

		dirs := []FsEntry{}
		files := []FsEntry{}

		for _, file := range osFiles {
			if file.IsDir() {
				dirs = append(dirs, FsEntry{
					Name: file.Name(),
				})
			} else {
				files = append(files, FsEntry{
					Name: file.Name(),
				})
			}

		}

		payload := Resp{
			Files: files,
			Dirs:  dirs,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})

}
