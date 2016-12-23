package api

import (
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
}
