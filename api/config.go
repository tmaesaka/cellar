package api

// config holds the server settings. These settings could come from
// the command-line and/or a configuration file.
type config struct {
	Port    int    `json:"port"`    // TCP port number that the server listens on
	DataDir string `json:"datadir"` // Path to the Cellar data directory
}

func NewConfig() *config {
	return &config{}
}
