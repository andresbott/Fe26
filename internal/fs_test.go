package internal

import (
	"encoding/json"
	"github.com/andresbott/go-carbon/libs/mock"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFsHandler(t *testing.T) {
	tcs := []struct {
		name       string
		req        func() (*http.Request, error)
		expecErr   string
		expectCode int
	}{
		{
			name: "successful request",
			req: func() (*http.Request, error) {

				req, err := http.NewRequest("GET", "/", nil)
				if err != nil {
					return nil, err
				}

				return req, nil
			},
			expectCode: http.StatusOK,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			req, err := tc.req()
			if err != nil {
				t.Fatal(err)
			}
			fs := mock.AferoHttpFs()

			handler := FileServer(fs)
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
			want := "banana"
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("unexpected value (-got +want)\n%s", diff)
			}

			//if tc.expecErr != "" {
			//	if status := recorder.Code; status != tc.expectCode {
			//		t.Errorf("handler returned wrong status code: got %v want %v",
			//			status, tc.expectCode)
			//	}
			//	respText, err := io.ReadAll(recorder.Body)
			//	if err != nil {
			//		t.Fatal(err)
			//	}
			//	got := strings.TrimSuffix(string(respText), "\n")
			//	if got != tc.expecErr {
			//		t.Errorf("unexpecter error message: got \"%s\" want \"%v\"",
			//			got, tc.expecErr)
			//	}
			//
			//} else {
			//
			//	if status := recorder.Code; status != tc.expectCode {
			//		t.Errorf("handler returned wrong status code: got %v want %v",
			//			status, tc.expectCode)
			//	}
			//
			//	got := localTaskList{}
			//	err = json.NewDecoder(recorder.Body).Decode(&got)
			//	if err != nil {
			//		t.Fatal(err)
			//	}
			//	if diff := cmp.Diff(got, tc.expect, cmpopts.IgnoreFields(localTaskOutput{}, "Id")); diff != "" {
			//		t.Errorf("unexpected value (-got +want)\n%s", diff)
			//	}
			//
			//}

		})
	}
}
