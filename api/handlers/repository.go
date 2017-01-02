package handlers

import (
	"net/http"

	"github.com/husobee/vestigo"
	"github.com/tmaesaka/cellar/config"
)

func IndexRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all repositories"))
	}
}

func ShowRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("showing " + vestigo.Param(r, "id")))
	}
}

func CreateRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("creating a repository"))
	}
}

func UpdateRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("updating " + vestigo.Param(r, "id")))
	}
}

func DestroyRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("destroying " + vestigo.Param(r, "id")))
	}
}
