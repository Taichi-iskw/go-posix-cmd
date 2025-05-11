package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	input := "hello\nworld\n"
	r:= strings.NewReader(input)
	var b bytes.Buffer

	err := Cat(r, &b)
	if err != nil {
		t.Fatalf("failed to cat: %v", err)
	}

	if b.String() != input {
		t.Errorf("expected %q, got %q", input, b.String())
	}
}