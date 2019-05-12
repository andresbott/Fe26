package f

import (
	"regexp"
	"strings"
	"errors"
	"path/filepath"
	"os"
)

func Sanitize( in string) (string,error){
	re := regexp.MustCompile(`/+`)
	s := re.ReplaceAllString(in, "/")

	r:=""
	if strings.HasPrefix(in,"/"){
		r = "/"
	}

	pathParts := strings.Split(s, "/")
	var newPath []string
	for _, item := range pathParts {
		if item == "."{
			// do nothing
		}else if item == ".."{
			if len(newPath) > 0 {
				newPath = newPath[:len(newPath)-1]
			}else {
				return "",errors.New("Reached top level, unable to go furthter up.")
			}
		}else {
			if item != ""{
				newPath = append(newPath, item)
			}

		}
	}

	r = r+strings.Join(newPath,"/")
	return r,nil
}

type FileTypeExtInfo struct {
	MimeType string
	Type string
}

func FileTypeFromExtension(ext string) (FileTypeExtInfo) {

	var r FileTypeExtInfo

	switch strings.ToLower(ext) {
	// Images
		case "jpeg":
			r = FileTypeExtInfo{
				MimeType: "image/jpeg",
				Type: "image",
			}
		case "jpg":
			r = FileTypeExtInfo{
				MimeType: "image/jpeg",
				Type: "image",
			}
		case "gif":
			r = FileTypeExtInfo{
				MimeType: "image/gif",
				Type: "image",
			}
		case "bmp":
			r = FileTypeExtInfo{
				MimeType: "image/bmp",
				Type: "image",
			}
		case "png":
			r = FileTypeExtInfo{
				MimeType: "image/png",
				Type: "image",
			}
		case "svg":
			r = FileTypeExtInfo{
				MimeType: "image/svg+xml",
				Type: "image",
			}
		case "webp":
			r = FileTypeExtInfo{
				MimeType: "image/webp",
				Type: "image",
			}
		case "tiff":
			r = FileTypeExtInfo{
				MimeType: "image/tiff",
				Type: "image",
			}
	// audio
		case "mp3":
			r = FileTypeExtInfo{
				MimeType: "audio/mpeg",
				Type: "audio",
			}
		case "ogg":
			r = FileTypeExtInfo{
				MimeType: "	audio/ogg",
				Type: "audio",
			}
		case "mid":
		case "midi":
			r = FileTypeExtInfo{
				MimeType: "	audio/midi",
				Type: "audio",
			}
		case "aac":
			r = FileTypeExtInfo{
				MimeType: "	audio/aac",
				Type: "audio",
			}
		case "wav":
			r = FileTypeExtInfo{
				MimeType: "	audio/wav",
				Type: "audio",
			}
		case "weba":
			r = FileTypeExtInfo{
				MimeType: "	audio/webm",
				Type: "audio",
			}
	// video
		case "mpeg":
			r = FileTypeExtInfo{
				MimeType: "video/mpeg",
				Type: "video",
			}
		case "mpg":
			r = FileTypeExtInfo{
				MimeType: "video/mpeg",
				Type: "video",
			}
		case "ogv":
			r = FileTypeExtInfo{
				MimeType: "video/ogg",
				Type: "video",
			}
		case "webm":
			r = FileTypeExtInfo{
				MimeType: "video/webm",
				Type: "video",
			}
		case "avi":
			r = FileTypeExtInfo{
				MimeType: "video/x-msvideo",
				Type: "video",
			}
		case "mp4":
			r = FileTypeExtInfo{
				MimeType: "video/mp4",
				Type: "video",
			}
		case "ts":
			r = FileTypeExtInfo{
				MimeType: "video/MP2T",
				Type: "video",
			}
		case "mkv":
			r = FileTypeExtInfo{
				MimeType: "video/x-matroska",
				Type: "video",
			}
	// File compression
		case "bz":
		case "tar.bz":
			r = FileTypeExtInfo{
				MimeType: "application/x-bzip",
				Type: "zip",
			}
		case "bz2":
			r = FileTypeExtInfo{
				MimeType: "application/x-bzip2",
				Type: "zip",
			}
		case "rar":
			r = FileTypeExtInfo{
				MimeType: "application/x-rar-compressed",
				Type: "zip",
			}
		case "tar":
			r = FileTypeExtInfo{
				MimeType: "application/x-tar",
				Type: "zip",
			}
		case "zip":
			r = FileTypeExtInfo{
				MimeType: "application/zip",
				Type: "zip",
			}
		case "7z":
			r = FileTypeExtInfo{
				MimeType: "application/x-7z-compressed",
				Type: "zip",
			}
	// text Files
		case "txt":
			r = FileTypeExtInfo{
				MimeType: "text/plain",
				Type: "text",
			}
		case "css":
			r = FileTypeExtInfo{
				MimeType: "text/css",
				Type: "text",
			}
		case "csv":
			r = FileTypeExtInfo{
				MimeType: "text/csv",
				Type: "text",
			}
		case "html":
			r = FileTypeExtInfo{
				MimeType: "text/html",
				Type: "web",
			}
		case "js":
			r = FileTypeExtInfo{
				MimeType: "text/javascript",
				Type: "text",
			}
		case "odt":
			r = FileTypeExtInfo{
				MimeType: "application/vnd.oasis.opendocument.text",
				Type: "text",
			}
		case "rtf":
			r = FileTypeExtInfo{
				MimeType: "text/rtf",
				Type: "text",
			}
		case "xml":
			r = FileTypeExtInfo{
				MimeType: "text/xml",
				Type: "text",
			}
	// Other
		case "exe":
			r = FileTypeExtInfo{
				MimeType: "application/octet-stream",
				Type: "executable",
			}
		case "sh":
			r = FileTypeExtInfo{
				MimeType: "text/plain",
				Type: "executable",
			}
		case "bash":
			r = FileTypeExtInfo{
				MimeType: "text/plain",
				Type: "executable",
			}
		case "py":
			r = FileTypeExtInfo{
				MimeType: "text/plain",
				Type: "executable",
			}
		case "php":
			r = FileTypeExtInfo{
				MimeType: "text/plain",
				Type: "executable",
			}
		case "pdf":
			r = FileTypeExtInfo{
				MimeType: "application/pdf",
				Type: "pdf",
			}

		default:
			r = FileTypeExtInfo{
				MimeType: "application/octet-stream",
				Type: "file",
			}
		}
	return r
}

// return the absolute path of a path string
func GetAbsPath (path string) (string,error){
	abs,_ := filepath.Abs(path)
	_,err := os.Stat(abs)
	if err != nil {
		return "",errors.New("Path does not exist: "+abs)
	}
	return abs,nil
}

// return the absolute path directory of a path string
func GetAbsPathDir(path string) (string,error){
	abs,_ := filepath.Abs(path)
	FI,err := os.Stat(abs)
	if err != nil {
		return "",errors.New("Path does not exist: "+abs)
	}

	if ! FI.Mode().IsDir() {
		return "",errors.New("Not a directory: "+abs)
	}

	return abs,nil
}