package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/husobee/vestigo"
	"github.com/libgit2/git2go"
	"github.com/tmaesaka/cellar/config"
)

// CreateContentHandler creates a new file in the specified repository.
// I still need to do some research but in theory we should be able to
// write to the bare git repository without creating a working repo.
//
// 1) Create a blob based on the provided content (git hash-object)
// 2) Add the blob to repository index (git update-index)
// 3) Write the index to a tree (git write-tree)
// 4) Commit the tree (git commit-tree)
// 5) Update the ref to point at the latest commit (git update-ref)
func CreateContentHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repoName := vestigo.Param(r, "name")
		repoPath := repoPath(cfg.DataDir, repoName)

		content := r.FormValue("content")
		contentPath := vestigo.Param(r, "_name")

		_, err := git.OpenRepository(repoPath)
		if err != nil {
			NotFound(w)
			return
		}

		if len(content) == 0 {
			BadRequest(w, InvalidRequestError, "content parameter required")
			return
		}

		_, err = base64.StdEncoding.DecodeString(content)
		if err != nil {
			BadRequest(w, InvalidRequestError, "content must be base64 encoded")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(contentPath))
	}
}
