package handlers

import (
	"fmt"
	"net/http"

	"github.com/husobee/vestigo"
	"github.com/libgit2/git2go"
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

// CreateRepositoryHandler provisions a bare git repository under the datadir
// directory. Relevant validation is also executed.
func CreateRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")

		if len(name) == 0 {
			renderError(w, ErrorInvalidRequest, "name parameter required")
			return
		}

		rpath := repoPath(cfg.DataDir, name)
		gitRepo, _ := git.OpenRepository(rpath)

		if gitRepo != nil {
			errStr := fmt.Sprintf("respoitory %s already exists", name)
			renderError(w, ErrorApi, errStr)
			return
		}

		bareRepo := true
		_, err := git.InitRepository(rpath, bareRepo)

		if err != nil {
			renderError(w, ErrorApi, "failed to init repository")
			return
		}

		w.Write([]byte("success"))
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
