package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTouchCreatesFile(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "newfile.txt")

	err := Touch(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("file was not created: %v", err)
	}

	if info.Size() != 0 {
		t.Errorf("expected empty file, got size %d", info.Size())
	}
}

func TestTouchUpdatesTimestamp(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "file.txt")

	err := os.WriteFile(path, []byte("data"), 0644)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	oldTime := time.Now().Add(-time.Hour)
	err = os.Chtimes(path, oldTime, oldTime)
	if err != nil {
		t.Fatalf("failed to set old time: %v", err)
	}

	err = Touch(path)
	if err != nil {
		t.Fatalf("touch failed: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat failed: %v", err)
	}

	if info.ModTime().Before(oldTime) {
		t.Errorf("mod time was not updated")
	}
}
