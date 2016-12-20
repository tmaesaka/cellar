package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Run checks if the provided configuration is sufficient to run the
// Cellar daemon. If successful, a Web API server will be started.
func Run(config *config) error {
	fmt.Fprintf(os.Stderr, "Starting cellard... listening on port %d\n", config.Port)

	addr := fmt.Sprintf(":%d", config.Port)

	// Start an empty http server until we decide on handler strategy.
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
