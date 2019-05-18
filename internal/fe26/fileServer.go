package fe26

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"errors"
	"path/filepath"
	"path"
	"os"
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


