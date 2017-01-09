package config

// ApiConfig holds the server settings. These settings could come from
// the command-line and/or a configuration file.
type ApiConfig struct {
	Port        int    `json:"port"`         // TCP port number that the server listens on
	DataDir     string `json:"datadir"`      // Path to the Cellar data directory
	Verbose     bool   `json:"verbose"`      // Server is verbose
	VeryVerbose bool   `json:"very_verbose"` // Server is very verbose
}

// NewConfig returns a new config for server configuration.
func NewApiConfig() *ApiConfig {
	return &ApiConfig{}
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
