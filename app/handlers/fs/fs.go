package fs

import (
	"encoding/json"
	"errors"
	"io/fs"
	"net/http"
	"path"
	"sort"
	"strings"
	"time"
)

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at root.
//
// this is a wrapper around the http.FileServer but generating
// json responses instead of html
//
// To use the operating system's file system implementation,
// use http.Dir:
//
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// To use an fs.FS implementation, use http.FS to convert it:
//
//	http.Handle("/", http.FileServer(http.FS(fsys)))
//
// To use afero use:
// http.Handle("/", http.FileServer(afero.NewHttpFs(afero.FS)))
func FileServer(root http.FileSystem, prefix string) http.Handler {
	return &fileHandler{root, prefix}
}

type fileHandler struct {
	root   http.FileSystem
	prefix string
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	upath = strings.TrimPrefix(upath, f.prefix)
	serveFile(w, r, f.root, path.Clean(upath))
}

// name is '/'-separated, not filepath.Separator.
func serveFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem, name string) {

	f, err := fs.Open(name)
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}

	// redirect to canonical path: / at end of directory url
	// r.URL.Path always begins with /
	pathUrl := r.URL.Path
	if d.IsDir() {
		if pathUrl[len(pathUrl)-1] != '/' {
			localRedirect(w, r, path.Base(pathUrl)+"/")
			return
		}
	} else {
		if pathUrl[len(pathUrl)-1] == '/' {
			localRedirect(w, r, "../"+path.Base(pathUrl))
			return
		}
	}

	if d.IsDir() {
		if checkIfModifiedSince(r, d.ModTime()) == condFalse {
			writeNotModified(w)
			return
		}
		setLastModified(w, d.ModTime())
		dirList(w, r, f)
		return
	}
	http.ServeContent(w, r, d.Name(), d.ModTime(), f)
}

type anyDirs interface {
	len() int
	name(i int) string
	isDir(i int) bool
	size(i int) int64
	modTime(i int) time.Time
}

type fileInfoDirs []fs.FileInfo

func (d fileInfoDirs) len() int                { return len(d) }
func (d fileInfoDirs) isDir(i int) bool        { return d[i].IsDir() }
func (d fileInfoDirs) name(i int) string       { return d[i].Name() }
func (d fileInfoDirs) size(i int) int64        { return d[i].Size() }
func (d fileInfoDirs) modTime(i int) time.Time { return d[i].ModTime() }

type fileInfoPayload struct {
	Name    string
	Size    int64
	ModTime time.Time
	IsDir   bool
}
type dirListPayload struct {
	Items []fileInfoPayload
}

func dirList(w http.ResponseWriter, r *http.Request, f http.File) {
	// here we defer fom http.fs because we want to return more detailed information
	// TODO: maybe expose a way to gather the information the user needs, e.g. with a callback ?
	var dirs anyDirs
	var list fileInfoDirs
	list, err := f.Readdir(-1)
	dirs = list

	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs.name(i) < dirs.name(j) })
	out := []fileInfoPayload{}
	for i, n := 0, dirs.len(); i < n; i++ {
		out = append(out, fileInfoPayload{
			Name:    dirs.name(i),
			Size:    dirs.size(i),
			ModTime: dirs.modTime(i),
			IsDir:   dirs.isDir(i),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	payload := dirListPayload{
		Items: out,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error marshaling the payload", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

// localRedirect gives a Moved Permanently response.
// It does not convert relative paths to absolute paths like Redirect does.
func localRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}

// toHTTPError returns a non-specific HTTP error message and status code
// for a given non-nil error value. It's important that toHTTPError does not
// actually return err.Error(), since msg and httpStatus are returned to users,
// and historically Go's ServeContent always returned just "404 Not Found" for
// all errors. We don't want to start leaking information in error messages.
func toHTTPError(err error) (msg string, httpStatus int) {
	if errors.Is(err, fs.ErrNotExist) {
		return "404 page not found", http.StatusNotFound
	}
	if errors.Is(err, fs.ErrPermission) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}

func setLastModified(w http.ResponseWriter, modtime time.Time) {
	if !isZeroTime(modtime) {
		w.Header().Set("Last-Modified", modtime.UTC().Format(TimeFormat))
	}
}

func writeNotModified(w http.ResponseWriter) {
	// RFC 7232 section 4.1:
	// a sender SHOULD NOT generate representation metadata other than the
	// above listed fields unless said metadata exists for the purpose of
	// guiding cache updates (e.g., Last-Modified might be useful if the
	// response does not have an ETag field).
	h := w.Header()
	delete(h, "Content-Type")
	delete(h, "Content-Length")
	delete(h, "Content-Encoding")
	if h.Get("Etag") != "" {
		delete(h, "Last-Modified")
	}
	w.WriteHeader(http.StatusNotModified)
}

// condResult is the result of an HTTP request precondition check.
// See https://tools.ietf.org/html/rfc7232 section 3.
type condResult int

const (
	condNone condResult = iota
	condTrue
	condFalse
)

func checkIfModifiedSince(r *http.Request, modtime time.Time) condResult {
	if r.Method != "GET" && r.Method != "HEAD" {
		return condNone
	}
	ims := r.Header.Get("If-Modified-Since")
	if ims == "" || isZeroTime(modtime) {
		return condNone
	}
	t, err := ParseTime(ims)
	if err != nil {
		return condNone
	}
	// The Last-Modified header truncates sub-second precision so
	// the modtime needs to be truncated too.
	modtime = modtime.Truncate(time.Second)
	if ret := modtime.Compare(t); ret <= 0 {
		return condFalse
	}
	return condTrue
}

var unixEpochTime = time.Unix(0, 0)

// isZeroTime reports whether t is obviously unspecified (either zero or Unix()=0).
func isZeroTime(t time.Time) bool {
	return t.IsZero() || t.Equal(unixEpochTime)
}

// TimeFormat is the time format to use when generating times in HTTP
// headers. It is like time.RFC1123 but hard-codes GMT as the time
// zone. The time being formatted must be in UTC for Format to
// generate the correct format.
//
// For parsing this time format, see ParseTime.
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

var timeFormats = []string{
	TimeFormat,
	time.RFC850,
	time.ANSIC,
}

// ParseTime parses a time header (such as the Date: header),
// trying each of the three formats allowed by HTTP/1.1:
// TimeFormat, time.RFC850, and time.ANSIC.
func ParseTime(text string) (t time.Time, err error) {
	for _, layout := range timeFormats {
		t, err = time.Parse(layout, text)
		if err == nil {
			return
		}
	}
	return
}
