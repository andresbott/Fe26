package fileserver

import (
	"github.com/andresbott/go-carbon/libs/mock"
	"net/http"
	"net/http/httptest"
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
