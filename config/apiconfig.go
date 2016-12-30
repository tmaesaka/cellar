package config

import (
	"strings"
)

// ApiConfig holds the server settings. These settings could come from
// the command-line and/or a configuration file.
type ApiConfig struct {
	Port        int    `json:"port"`         // TCP port number that the server listens on
	DataDir     string `json:"datadir"`      // Path to the Cellar data directory
	Verbose     bool   `json:"verbose"`      // Server is verbose
	VeryVerbose bool   `json:"very_verbose"` // Server is very verbose
}

// configValidationError is a custom error type for representing
// validation related errors.
type configValidationError struct {
	InvalidFields map[string]string // Maps a field to its descriptive error message
}

// NewConfig returns a new config for server configuration.
func NewApiConfig() *ApiConfig {
	return &ApiConfig{}
}

// NewConfigValidationError returns a new configValidationError.
func NewConfigValidationError() *configValidationError {
	err := configValidationError{}
	err.InvalidFields = make(map[string]string)
	return &err
}

// Validate validates a given config.
func (cfg *ApiConfig) Validate() error {
	rv := NewConfigValidationError()

	if cfg.Port == 0 {
		rv.InvalidFields["port"] = "tcp port number is invalid"
	}
	if cfg.DataDir == "" {
		rv.InvalidFields["datadir"] = "datadir can't be blank"
	}

	if len(rv.InvalidFields) > 0 {
		return rv
	} else {
		return nil
	}
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
