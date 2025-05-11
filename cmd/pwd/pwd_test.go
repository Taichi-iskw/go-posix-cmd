package main

import (
	"testing"
)

func TestGetWorkingDir(t *testing.T) {
	dir, err := GetWorkingDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dir == "" {
		t.Fatal("empty working directory returned")
	}
}
