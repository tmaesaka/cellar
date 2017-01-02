package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/husobee/vestigo"
	"github.com/tmaesaka/cellar/config"
)

var (
	cfg    = config.NewApiConfig()
	repoId = "project_x"
)

func TestIndexRepositoryHandler(t *testing.T) {
	handler := IndexRepositoryHandler(cfg)
	req, _ := http.NewRequest("GET", "/repositories", nil)

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
	router := vestigo.NewRouter()
	router.Get("/repositories/:id", ShowRepositoryHandler(cfg))

	path := "/repositories/" + repoId
	req, _ := http.NewRequest("GET", path, nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	if recorder.Body.String() != "showing "+repoId {
		t.Error("Unexpected response body")
	}
}

func TestCreateRepositoryHandler(t *testing.T) {
	router := vestigo.NewRouter()
	router.Post("/repositories", CreateRepositoryHandler(cfg))

	badReq, _ := http.NewRequest("POST", "/repositories", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, badReq)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Exepected status code 400; got %d", recorder.Code)
	}

	err := decodeError(recorder.Body.String())

	if err.ErrorType != ErrorInvalidRequest {
		t.Errorf("Expected invalid_request_error; got %s", err.ErrorType)
	}

	if err.Message != "name parameter required" {
		t.Errorf("Unexpected response; got %s", err.Message)
	}
}

func TestUpdateRepositoryHandler(t *testing.T) {
	router := vestigo.NewRouter()
	router.Put("/repositories/:id", UpdateRepositoryHandler(cfg))

	path := "/repositories/" + repoId
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
	router := vestigo.NewRouter()
	router.Delete("/repositories/:id", DestroyRepositoryHandler(cfg))

	path := "/repositories/" + repoId
	req, _ := http.NewRequest("DELETE", path, nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	if recorder.Body.String() != "destroying "+repoId {
		t.Error("Unexpected response body")
	}
}
