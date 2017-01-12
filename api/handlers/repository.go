package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/husobee/vestigo"
	"github.com/libgit2/git2go"
	"github.com/tmaesaka/cellar/config"
)

// Repository type holds information about a repository.
type Repository struct {
	Name string `json:"name"` // Unique name of the repository
}

// IndexRepositoryHandler generates a list of existing repositories.
func IndexRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repos := make([]Repository, 0)

		files, err := ioutil.ReadDir(path.Join(cfg.DataDir, "repos"))

		if err != nil {
			BadRequest(w, ApiError, err.Error())
			return
		}

		for _, f := range files {
			repoPath := repoPath(cfg.DataDir, f.Name())
			_, err := git.OpenRepositoryExtended(repoPath, git.RepositoryOpenNoSearch, "")

			if err != nil {
				if cfg.Verbose {
					log.Printf("%s is not a git repository", repoPath)
				}
				continue
			}

			repos = append(repos, Repository{Name: f.Name()})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repos)
	}
}

// ShowRepositoryHandler looks up the requested git repository.
func ShowRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := vestigo.Param(r, "name")
		rpath := repoPath(cfg.DataDir, name)
		_, err := git.OpenRepository(rpath)

		if err != nil {
			NotFound(w)
			return
		}

		repo := Repository{Name: name}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repo)
	}
}

// CreateRepositoryHandler provisions a bare git repository under the datadir
// directory. Relevant validation is also executed.
func CreateRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")

		if len(name) == 0 {
			BadRequest(w, InvalidRequestError, "name parameter required")
			return
		}

		rpath := repoPath(cfg.DataDir, name)
		gitRepo, _ := git.OpenRepository(rpath)

		if gitRepo != nil {
			errStr := fmt.Sprintf("respoitory %s already exists", name)
			BadRequest(w, ApiError, errStr)
			return
		}

		bareRepo := true
		_, err := git.InitRepository(rpath, bareRepo)

		if err != nil {
			BadRequest(w, ApiError, "failed to init repository")
			return
		}

		repo := Repository{Name: name}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repo)
	}
}

func UpdateRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("unimplemented"))
	}
}

// DestroyRepositoryHandler destroys the specified repository.
func DestroyRepositoryHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := vestigo.Param(r, "name")
		rpath := repoPath(cfg.DataDir, name)

		if _, err := os.Stat(rpath); os.IsNotExist(err) {
			NotFound(w)
			return
		}

		if err := os.RemoveAll(rpath); err != nil {
			BadRequest(w, ApiError, "failed to destroy repository")
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
