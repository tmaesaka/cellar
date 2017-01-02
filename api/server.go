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
func Run(cfg *config.ApiConfig) error {
	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := ensureDataDirPresence(cfg.DataDir); err != nil {
		log.Fatalf("Invalid datadir: %v", err)
	}

	router := vestigo.NewRouter()
	router.Get("/config", handlers.IndexConfigHandler(cfg))
	router.Get("/repositories", handlers.IndexRepositoryHandler(cfg))
	router.Get("/repositories/:id", handlers.ShowRepositoryHandler(cfg))
	router.Post("/repositories", handlers.CreateRepositoryHandler(cfg))
	router.Put("/repositories/:id", handlers.UpdateRepositoryHandler(cfg))
	router.Delete("/repositories/:id", handlers.DestroyRepositoryHandler(cfg))

	fmt.Fprintf(os.Stderr, "Starting cellard... listening on port %d\n", cfg.Port)

	addr := fmt.Sprintf(":%d", cfg.Port)

	// Start an empty http server until we decide on handler strategy.
	log.Fatal(http.ListenAndServe(addr, router))

	return nil
}
