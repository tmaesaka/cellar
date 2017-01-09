package config

import (
	"testing"
)

func TestNewApiConfig(t *testing.T) {
	if NewApiConfig() == nil {
		t.Error("Failed to create config")
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
