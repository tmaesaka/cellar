package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/husobee/vestigo"
	"github.com/libgit2/git2go"
	"github.com/tmaesaka/cellar/config"
)

var (
	cfg     = config.NewApiConfig()
	repoId  = "project_x"
	testDir = "/tmp/_cellar_test"
)

func TestIndexRepositoryHandler(t *testing.T) {
	handler := IndexRepositoryHandler(cfg)
	req, _ := http.NewRequest("GET", "/repos", nil)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	if recorder.Body.String() != "all repositories" {
		t.Error("Unexpected response body")
	}
}

func TestShowRepositoryHandler(t *testing.T) {
	cfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Get("/repos/:id", ShowRepositoryHandler(cfg))

	path := path.Join("/repos", repoId)
	req, _ := http.NewRequest("GET", path, nil)

	t.Run("non-existing repository", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusNotFound {
			t.Errorf("Exepected status code 404; got %d", recorder.Code)
		}

		if len(recorder.Body.String()) != 0 {
			t.Errorf("Exepected an empty body")
		}
	})

	t.Run("existing repository", func(t *testing.T) {
		rpath := repoPath(cfg.DataDir, repoId)
		git.InitRepository(rpath, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Exepected status code 200; got %d", recorder.Code)
		}

		var resp Repository

		if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
			t.Error(err)
		}

		if resp.Name != repoId {
			t.Errorf("Expected %s; got %s", repoId, resp.Name)
		}
	})

	cleanup()
}

func TestCreateRepositoryHandler(t *testing.T) {
	cfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Post("/repos", CreateRepositoryHandler(cfg))

	t.Run("missing name param", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/repos", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Exepected status code 400; got %d", recorder.Code)
		}

		err := decodeError(recorder.Body.String())

		if err.ErrorType != InvalidRequestError {
			t.Errorf("Expected invalid_request_error; got %s", err.ErrorType)
		}

		if err.Message != "name parameter required" {
			t.Errorf("Unexpected response; got %s", err.Message)
		}
	})

	t.Run("name=<repoId>", func(t *testing.T) {
		params := url.Values{}
		params.Set("name", repoId)
		encodedParams := strings.NewReader(params.Encode())

		req, _ := http.NewRequest("POST", "/repos", encodedParams)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Exepected status code 200; got %d", recorder.Code)
		}

		var resp Repository

		if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
			t.Error(err)
		}

		if resp.Name != repoId {
			t.Errorf("Expected %s; got %s", repoId, resp.Name)
		}
	})

	cleanup()
}

func TestUpdateRepositoryHandler(t *testing.T) {
	router := vestigo.NewRouter()
	router.Put("/repos/:id", UpdateRepositoryHandler(cfg))

	path := path.Join("/repos", repoId)
	req, _ := http.NewRequest("PUT", path, nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	if recorder.Body.String() != "updating "+repoId {
		t.Error("Unexpected response body")
	}
}

func TestDestroyRepositoryHandler(t *testing.T) {
	cfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Delete("/repos/:id", DestroyRepositoryHandler(cfg))

	path := path.Join("/repos", repoId)

	t.Run("deleting a non-existing repo", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", path, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusNotFound {
			t.Errorf("Exepected status code 404; got %d", recorder.Code)
		}

		if len(recorder.Body.String()) != 0 {
			t.Errorf("Exepected an empty body")
		}
	})

	t.Run("deleting an existing repo", func(t *testing.T) {
		rpath := repoPath(cfg.DataDir, repoId)
		git.InitRepository(rpath, true)

		req, _ := http.NewRequest("DELETE", path, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Exepected status code 200; got %d", recorder.Code)
		}
	})

	cleanup()
}

func cleanup() {
	if err := os.RemoveAll(cfg.DataDir); err != nil {
		panic("failed to cleanup")
	}
}
