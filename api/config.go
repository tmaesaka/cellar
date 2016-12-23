package api

// config holds the server settings. These settings could come from
// the command-line and/or a configuration file.
type config struct {
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
func NewConfig() *config {
	return &config{}
}

// NewConfigValidationError returns a new configValidationError.
func NewConfigValidationError() *configValidationError {
	err := configValidationError{}
	err.InvalidFields = make(map[string]string)
	return &err
}
