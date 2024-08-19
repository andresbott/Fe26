package fileserver

import (
	"net/http"
	"strings"
)

func (f *fileHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	upath = strings.TrimPrefix(upath, f.prefix)
	if upath == "" || upath == "/" {
		http.Error(w, "refusing to delete the root directory", http.StatusBadRequest)
	}
	err := f.fs.RemoveAll(upath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
