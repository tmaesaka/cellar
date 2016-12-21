package api

import (
	"os"
	"testing"
)

func TestEnsureDataDirPresence(t *testing.T) {
	path := "/tmp/cellar_test"

	if err := ensureDataDirPresence(path); err != nil {
		t.Errorf("Expected mkdir %s to succeed", path)
		return
	}

	if err := os.Remove(path); err != nil {
		t.Error(err)
	}
}
