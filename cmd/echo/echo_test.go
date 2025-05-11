package main

import (
	"bytes"
	"testing"
)

func TestEchoWithNewline(t *testing.T) {
	var buf bytes.Buffer
	err := Echo(&buf, []string{"hello", "world"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "hello world\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestEchoWithoutNewline(t *testing.T) {
	var buf bytes.Buffer
	err := Echo(&buf, []string{"hello", "world"}, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "hello world"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
