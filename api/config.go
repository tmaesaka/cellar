package api

// config holds the server settings. These settings could come from
// the command-line and/or a configuration file.
type config struct {
	Port        int    `json:"port"`         // TCP port number that the server listens on
	DataDir     string `json:"datadir"`      // Path to the Cellar data directory
	Verbose     bool   `json:"verbose"`      // Server is verbose
	VeryVerbose bool   `json:"very_verbose"` // Server is very verbose
}

func NewConfig() *config {
	return &config{}
}
