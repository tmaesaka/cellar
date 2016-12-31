package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tmaesaka/cellar/config"
)

func TestConfigIndexHandler(t *testing.T) {
	cfg := config.NewApiConfig()
	handler := ConfigIndexHandler(cfg)
	req, _ := http.NewRequest("GET", "/config", nil)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Exepected status code 200; got %d", recorder.Code)
	}

	contentType := recorder.HeaderMap["Content-Type"][0]

	if contentType != "application/json" {
		t.Errorf("Expected Content-Type to be application/json; got %v", contentType)
	}

	var resp config.ApiConfig

	if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
}
