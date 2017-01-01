package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/husobee/vestigo"
	"github.com/tmaesaka/cellar/api/handlers"
	"github.com/tmaesaka/cellar/config"
)

func ensureDataDirPresence(datadir string) error {
	if _, err := os.Stat(datadir); os.IsNotExist(err) {
		if err = os.MkdirAll(datadir, 0755); err != nil {
			return err
		}
	}

	return nil
}

// Run checks if the provided configuration is sufficient to run the
// Cellar daemon. If successful, a Web API server will be started.
func Run(config *config.ApiConfig) error {
	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := ensureDataDirPresence(config.DataDir); err != nil {
		log.Fatalf("Invalid datadir: %v", err)
	}

	router := vestigo.NewRouter()
	router.Get("/config", handlers.IndexConfigHandler(config))

	fmt.Fprintf(os.Stderr, "Starting cellard... listening on port %d\n", config.Port)

	addr := fmt.Sprintf(":%d", config.Port)

	// Start an empty http server until we decide on handler strategy.
	log.Fatal(http.ListenAndServe(addr, router))

	return nil
}
