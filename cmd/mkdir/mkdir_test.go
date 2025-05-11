package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMkdirWithoutParents(t *testing.T) {
	tmp := t.TempDir()
	dir := filepath.Join(tmp, "subdir")
	
	err := Mkdir(dir, false)
	if err != nil {
		t.Fatalf("failed to create directory: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("failed to stat directory: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("created file is not a directory")
	}
}

func TestMkdirWithParents(t *testing.T) {
	tmp := t.TempDir()
	dir := filepath.Join(tmp, "parent", "subdir")
	
	err := Mkdir(dir, true)
	if err != nil {
		t.Fatalf("failed to create directory: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("failed to stat directory: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("created file is not a directory")
	}
}
