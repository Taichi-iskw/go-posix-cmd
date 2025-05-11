package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	tmp := t.TempDir()
	file := filepath.Join(tmp, "file.txt")

	err := os.WriteFile(file, []byte("hello"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	if err := Rm(file, false, false); err != nil {
		t.Fatalf("remove failed: %v", err)
	}

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		t.Errorf("file still exists after removal")
	}
}

func TestRemoveDirRecursive(t *testing.T) {
	tmp := t.TempDir()
	dir := filepath.Join(tmp, "a", "b", "c")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		t.Fatalf("failed to create nested dirs: %v", err)
	}

	if err := Rm(filepath.Join(tmp, "a"), true, false); err != nil {
		t.Fatalf("recursive remove failed: %v", err)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Errorf("directory still exists after recursive removal")
	}
}

func TestRemoveNonExistentWithForce(t *testing.T) {
	err := Rm("/nonexistent/path/to/file", false, true)
	if err != nil {
		t.Errorf("expected no error with force, got: %v", err)
	}
}
