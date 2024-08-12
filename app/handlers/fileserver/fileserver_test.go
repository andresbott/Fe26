package fileserver

import (
	"encoding/json"
	"fmt"
	"github.com/andresbott/go-carbon/libs/mock"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

var _ = spew.Dump // prevent from being removed by the IDE

var sampleDirs = []string{
	"/media/photos",
	"/media/video",
	"/media/music",
	"/text/plain/books",
	"/text/pdf",
	"/tree/a/a_b/a_b_a",
	"/tree/a/a_c/a_c_a",
}

var SampleFiles = map[string]string{
	"/text/plain/file1.txt": "file1.txt",
	"/text/plain/file2.txt": "file2.txt",
}

func TestFsHandler(t *testing.T) {
	tcs := []struct {
		name       string
		prefix     string
		req        func() (*http.Request, error)
		want       dirContent
		expectCode int
	}{
		{
			name: "list root dir",
			req: func() (*http.Request, error) {

				req, err := http.NewRequest("GET", "/", nil)
				if err != nil {
					return nil, err
				}
				return req, nil
			},
			expectCode: http.StatusOK,
			want: dirContent{
				Dirs: []string{"media", "text", "tree"},
			},
		},
		{
			name:   "list root dir with request prefix",
			prefix: "/api/v0/",
			req: func() (*http.Request, error) {

				req, err := http.NewRequest("GET", "/api/v0/", nil)
				if err != nil {
					return nil, err
				}
				return req, nil
			},
			expectCode: http.StatusOK,
			want: dirContent{
				Dirs: []string{"media", "text", "tree"},
			},
		},
		{
			name: "list files",
			req: func() (*http.Request, error) {

				req, err := http.NewRequest("GET", "/text/plain/", nil)
				if err != nil {
					return nil, err
				}
				return req, nil
			},
			expectCode: http.StatusOK,
			want: dirContent{
				Dirs:  []string{"books"},
				Files: []string{"file1.txt", "file2.txt"},
			},
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

			got := map[string]any{}
			err = json.NewDecoder(recorder.Body).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}
			content, err := decodeResp(got)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(content, tc.want); diff != "" {
				t.Errorf("unexpected value (-got +want)\n%s", diff)
			}

		})
	}
}

type dirContent struct {
	Dirs  []string
	Files []string
}

func decodeResp(in map[string]any) (dirContent, error) {
	c := dirContent{}
	for k, _ := range in {
		if k == "Items" {
			for _, fileEntry := range in["Items"].([]any) {
				data := fileEntry.(map[string]any)

				name, ok := data["Name"].(string)
				if !ok {
					return c, fmt.Errorf("the file item does not contain the key Name")
				}
				isDir, ok := data["IsDir"].(bool)
				if !ok {
					return c, fmt.Errorf("the file item does not contain the key Name")
				}
				if isDir {
					c.Dirs = append(c.Dirs, name)
				} else {
					c.Files = append(c.Files, name)
				}
			}
			break
		}
	}
	return c, nil
}
