package api

import (
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {
	if NewConfig() == nil {
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
