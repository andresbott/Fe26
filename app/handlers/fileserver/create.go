package fileserver

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

// handlePut handles directory creations
func (f *fileHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	upath = strings.TrimPrefix(upath, f.prefix)

	err := f.fs.Mkdir(upath, 0750)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// handlePost handles file uploads
func (f *fileHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	upath = strings.TrimPrefix(upath, f.prefix)

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(err)
	}

	// key = "file" in the form
	for key, files := range r.MultipartForm.File {
		_ = key // todo, eventyally only alow certain keys?

		for _, file := range files {
			//spew.Dump(file.Filename)
			//fmt.Printf("Uploaded File: %+v\n", file.Filename)
			//fmt.Printf("File Size: %+v\n", file.Size)
			//fmt.Printf("MIME Header: %+v\n", file.Header)
			fileDest := filepath.Join(upath, file.Filename)

			// create the destination

			dst, err := f.fs.Create(fileDest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// open the upload file
			f, err := file.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = io.Copy(dst, f)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}
