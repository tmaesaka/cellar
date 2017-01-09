package config

import (
	"strings"
)

// configValidationError is a custom error type for representing
// validation related errors.
type configValidationError struct {
	InvalidFields map[string]string // Maps a field to its descriptive error message
}

// NewConfigValidationError returns a new configValidationError.
func NewConfigValidationError() *configValidationError {
	err := configValidationError{}
	err.InvalidFields = make(map[string]string)
	return &err
}

// Error returns a summarized context of the validation error.
// It is also needed to implement the Go "error" interface.
func (cve *configValidationError) Error() string {
	if len(cve.InvalidFields) == 0 {
		return ""
	}

	errStr := "Invalid configuration found in: "
	keys := make([]string, 0, len(cve.InvalidFields))

	for key := range cve.InvalidFields {
		keys = append(keys, key)
	}

	return errStr + strings.Join(keys, ", ")
}
