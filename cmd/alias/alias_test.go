package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestAddAlias(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test_aliases.json")
	aliasFile = tmpFile

	err := AddAlias("ll", "ls -la")
	if err != nil {
		t.Fatalf("AddAlias failed: %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read saved alias file: %v", err)
	}

	var aliases map[string]string
	err = json.Unmarshal(data, &aliases)
	if err != nil {
		t.Fatalf("failed to parse alias JSON: %v", err)
	}

	if aliases["ll"] != "ls -la" {
		t.Errorf("expected 'ls -la', got %q", aliases["ll"])
	}
}
