package config

import (
	"strings"
	"testing"
)

func TestNewApiConfig(t *testing.T) {
	if NewApiConfig() == nil {
		t.Error("Failed to create config")
	}
}

func TestNewConfigValidationError(t *testing.T) {
	err := NewConfigValidationError()

	if err == nil {
		t.Errorf("Failed to create configValidationError")
	}

	if err.InvalidFields == nil {
		t.Error("ErrorMap hasn't been allocated memory")
	}

	if err.Error() != "" {
		t.Errorf("Expected a blank string; got %s", err.Error())
	}

	err.InvalidFields["datadir"] = "datadir can't be blank"

	if !strings.Contains(err.Error(), "datadir") {
		t.Errorf("Invalid error string; expected to contain datadir")
	}
}

func TestConfigValidate(t *testing.T) {
	config := NewApiConfig()
	err := config.Validate()

	if err == nil {
		t.Fatal("Expected config validation to fail")
	}

	if cve, ok := err.(*configValidationError); ok {
		if _, ok := cve.InvalidFields["port"]; !ok {
			t.Error("Expected port to fail validation")
		}
		if _, ok := cve.InvalidFields["datadir"]; !ok {
			t.Error("Expected datadir to fail validation")
		}
	}
}
