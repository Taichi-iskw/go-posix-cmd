package main

import (
	"bytes"
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

func TestListAliases(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "aliases.json")
	aliasFile = tmpFile 

	aliases := map[string]string{
		"ll": "ls -la",
		"gs": "git status",
	}
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = ListAliases()
	if err != nil {
		t.Fatalf("ListAliases failed: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if !contains(output, "ll") || !contains(output, "ls -la") {
		t.Errorf("expected ll → ls -la in output, got:\n%s", output)
	}
	if !contains(output, "gs") || !contains(output, "git status") {
		t.Errorf("expected gs → git status in output, got:\n%s", output)
	}
}

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}