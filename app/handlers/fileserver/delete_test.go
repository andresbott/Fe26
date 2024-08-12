package fileserver

import (
	"fmt"
	"github.com/andresbott/go-carbon/libs/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDelete(t *testing.T) {
	tcs := []struct {
		name       string
		prefix     string
		req        func() (*http.Request, error)
		notExists  string
		expectCode int
	}{
		{
			name: "delete a file",
			req: func() (*http.Request, error) {
				req, _ := http.NewRequest("DELETE", "/text/plain/file2.txt", nil)
				return req, nil
			},
			notExists:  "/text/plain/file2.txt",
			expectCode: http.StatusOK,
		},
		{
			name: "delete a directory",
			req: func() (*http.Request, error) {
				req, _ := http.NewRequest("DELETE", "/text", nil)
				return req, nil
			},
			notExists:  "/text/plain",
			expectCode: http.StatusOK,
		},
		{
			name: "delete a non existent",
			req: func() (*http.Request, error) {
				req, _ := http.NewRequest("DELETE", "/banana", nil)
				return req, nil
			},
			notExists:  "/banana",
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

			_, err = fs.Stat(tc.notExists)
			if err != nil {
				if err.Error() != fmt.Sprintf("open %s: file does not exist", tc.notExists) {
					t.Errorf("unexpected error:%s", err)
				}
			} else {
				t.Errorf("expect stat to fail with file not found")
			}

		})
	}
}
