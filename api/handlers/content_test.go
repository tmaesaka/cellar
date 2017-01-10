package handlers

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/husobee/vestigo"
)

func TestCreateContentHandler(t *testing.T) {
	cfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Post("/repos/:id/contents/*", CreateContentHandler(cfg))

	testFile := "test.txt"
	reqPath := path.Join("/repos", repoId, "contents", testFile)

	req, _ := http.NewRequest("POST", reqPath, nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	if recorder.Body.String() != testFile {
		t.Errorf("Expected %s; got %s", testFile, recorder.Body.String())
	}
}
