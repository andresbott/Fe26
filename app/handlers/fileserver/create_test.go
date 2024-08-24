package fileserver

import (
	"bytes"
	"github.com/andresbott/go-carbon/libs/mock"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestCreateDir(t *testing.T) {
	tcs := []struct {
		name       string
		prefix     string
		req        func() (*http.Request, error)
		want       string
		expectCode int
	}{
		{
			name: "Create a dir",
			req: func() (*http.Request, error) {
				req, _ := http.NewRequest("PUT", "/text/plain/myDir", nil)
				return req, nil
			},
			want:       "/text/plain/myDir",
			expectCode: http.StatusOK,
		},
		{
			name:   "Create a dir with prefix",
			prefix: "/api/v0",
			req: func() (*http.Request, error) {
				req, _ := http.NewRequest("PUT", "/api/v0/text/plain/myDir", nil)
				return req, nil
			},
			want:       "/text/plain/myDir",
			expectCode: http.StatusOK,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			req, err := tc.req()
			if err != nil {
				t.Fatal(err)
			}
			fs := mock.AferoSample(sampleDirs, SampleFiles)

			handler := FileServer(fs, tc.prefix)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)

			if status := recorder.Code; status != tc.expectCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectCode)
			}

			fstat, err := fs.Stat(tc.want)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if fstat.Name() == "" {
				t.Errorf("expect dir to exitss")
			}
			if !fstat.IsDir() {
				t.Errorf("expect result to be a directory")
			}

		})
	}
}

func TestCreateFile(t *testing.T) {
	tcs := []struct {
		name       string
		prefix     string
		req        func() (*http.Request, error)
		want       string
		expectCode int
	}{
		{
			name: "upload single file",
			req: func() (*http.Request, error) {
				file := content{
					fname: "sample.txt",
					ftype: "demo",
					fdata: []byte("my content "),
				}
				buf, contentType, err := filesBuf(file)
				if err != nil {
					return nil, err
				}
				req, err := http.NewRequest("POST", "/text/plain/myDir", buf)
				if err != nil {
					return nil, err
				}
				req.Header.Add("Content-Type", contentType)

				return req, nil

			},
			want:       "/text/plain/myDir/sample.txt",
			expectCode: http.StatusOK,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			req, err := tc.req()
			if err != nil {
				t.Fatal(err)
			}
			fs := mock.AferoSample(sampleDirs, SampleFiles)

			handler := FileServer(fs, tc.prefix)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)

			if status := recorder.Code; status != tc.expectCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectCode)
			}

			fstat, err := fs.Stat(tc.want)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if fstat.Name() == "" {
				t.Errorf("expect file to exitss")
			}
			if fstat.IsDir() {
				t.Errorf("expect result to be a file")
			}

		})
	}
}

// content is a struct which contains a file's name, its type and its data.
type content struct {
	fname string
	ftype string
	fdata []byte
}

func filesBuf(files ...content) (*bytes.Buffer, string, error) {
	var (
		buf = new(bytes.Buffer)
		w   = multipart.NewWriter(buf)
	)
	w.FormDataContentType()

	for _, f := range files {
		part, err := w.CreateFormFile(f.ftype, filepath.Base(f.fname))
		if err != nil {
			return nil, "", err
		}

		_, err = part.Write(f.fdata)
		if err != nil {
			return nil, "", err
		}
	}

	err := w.Close()
	if err != nil {
		return nil, "", err
	}
	return buf, w.FormDataContentType(), nil

}
