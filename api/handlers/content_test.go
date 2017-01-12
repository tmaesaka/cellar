package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"

	"github.com/husobee/vestigo"
)

func TestCreateContentHandler(t *testing.T) {
	testCfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Post("/repos/:name/contents/*", CreateContentHandler(testCfg))

	testFile := "test.txt"
	reqPath := path.Join("/repos", testRepoName, "contents", testFile)

	t.Run("without content", func(t *testing.T) {
		req, _ := http.NewRequest("POST", reqPath, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Exepected status code 400; got %d", recorder.Code)
		}

		errResp := decodeError(recorder.Body.String())

		if errResp.Message != "content parameter required" {
			t.Errorf("Mismatched error message: %s", recorder.Body.String())
		}
	})

	t.Run("with base64 encoded content", func(t *testing.T) {
		params := url.Values{}
		params.Set("content", "Y2VsbGFy")
		encodedParams := strings.NewReader(params.Encode())

		req, _ := http.NewRequest("POST", reqPath, encodedParams)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Exepected status code 200; got %d", recorder.Code)
		}
	})

	t.Run("with non-base64 encoded content", func(t *testing.T) {
		params := url.Values{}
		params.Set("content", "cellar")
		encodedParams := strings.NewReader(params.Encode())

		req, _ := http.NewRequest("POST", reqPath, encodedParams)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Exepected status code 400; got %d", recorder.Code)
		}

		errResp := decodeError(recorder.Body.String())

		if errResp.Message != "content must be base64 encoded" {
			t.Errorf("Mismatched error message: %s", errResp.Message)
		}
	})
}
