package fileserver

import (
	"net/http"
	"strings"
)

func (f *fileHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	// todo implement, this should also allow creation of dirs
	w.WriteHeader(http.StatusOK)
}
