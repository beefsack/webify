package lib

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestServer_Cat(t *testing.T) {
	server := Server{
		Opts: Opts{
			Script: []string{"cat"},
			Addr:   "http://127.0.0.1",
		},
	}

	reqBody := []byte("blah")

	req := httptest.NewRequest("GET", "http://127.0.0.1/", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	if string(body) != string(reqBody) {
		t.Errorf("Expected body to be \"%s\" but got \"%s\"", reqBody, body)
	}
}
