package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tmaesaka/cellar/config"
)

// IndexConfigHandler encodes the current state of the server config to
// JSON and writes the result to the http connection.
func IndexConfigHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO(toru): Write a middleware that does this.
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(cfg); err != nil {
			log.Print(err)
		}
	}
}
