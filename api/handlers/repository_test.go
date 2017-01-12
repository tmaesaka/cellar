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
	testCfg      = config.NewApiConfig()
	testRepoName = "project_x"
	testDir      = "/tmp/_cellar_test"
)

func TestIndexRepositoryHandler(t *testing.T) {
	testCfg.DataDir = testDir
	handler := IndexRepositoryHandler(testCfg)
	req, _ := http.NewRequest("GET", "/repos", nil)

	t.Run("without repositories", func(t *testing.T) {
		// FIXME(toru): This is a dirty hack around cleanup()
		if err := os.MkdirAll(testCfg.DataDir+"/repos", 0755); err != nil {
			t.Error(err.Error())
		}

		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code 200; got %d", recorder.Code)
		}

		if recorder.Body.String() != "[]\n" {
			t.Errorf("Expected []; got %s", recorder.Body.String())
		}
	})

	t.Run("with repositories", func(t *testing.T) {
		repoNames := []string{"repo1", "repo2"}

		for _, name := range repoNames {
			git.InitRepository(repoPath(testCfg.DataDir, name), true)
		}

		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Exepected status code 200; got %d", recorder.Code)
		}

		var repos []Repository
		if err := json.NewDecoder(recorder.Body).Decode(&repos); err != nil {
			t.Error(err)
		}

		if len(repos) != 2 {
			t.Error("Expected 2 repositories; got %d", len(repos))
		}

		for idx, repo := range repos {
			if repo.Name != repoNames[idx] {
				t.Error("Expected %s; got %s", repo.Name, repoNames[idx])
			}
		}
	})

	cleanup()
}

func TestShowRepositoryHandler(t *testing.T) {
	testCfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Get("/repos/:name", ShowRepositoryHandler(testCfg))

	path := path.Join("/repos", testRepoName)
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
		rpath := repoPath(testCfg.DataDir, testRepoName)
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

		if resp.Name != testRepoName {
			t.Errorf("Expected %s; got %s", testRepoName, resp.Name)
		}
	})

	cleanup()
}

func TestCreateRepositoryHandler(t *testing.T) {
	testCfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Post("/repos", CreateRepositoryHandler(testCfg))

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

	t.Run("name=<repoName>", func(t *testing.T) {
		params := url.Values{}
		params.Set("name", testRepoName)
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

		if resp.Name != testRepoName {
			t.Errorf("Expected %s; got %s", testRepoName, resp.Name)
		}
	})

	cleanup()
}

func TestUpdateRepositoryHandler(t *testing.T) {
	router := vestigo.NewRouter()
	router.Put("/repos/:name", UpdateRepositoryHandler(testCfg))

	path := path.Join("/repos", testRepoName)
	req, _ := http.NewRequest("PUT", path, nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	if recorder.Body.String() != "unimplemented" {
		t.Error("Unexpected response body")
	}
}

func TestDestroyRepositoryHandler(t *testing.T) {
	testCfg.DataDir = testDir
	router := vestigo.NewRouter()
	router.Delete("/repos/:name", DestroyRepositoryHandler(testCfg))

	path := path.Join("/repos", testRepoName)

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
		rpath := repoPath(testCfg.DataDir, testRepoName)
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
	if err := os.RemoveAll(testCfg.DataDir); err != nil {
		panic("failed to cleanup")
	}
}
