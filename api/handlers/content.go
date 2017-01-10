package handlers

import (
	"net/http"

	"github.com/husobee/vestigo"
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
		path := vestigo.Param(r, "_name")

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(path))
	}
}
