package main

import (
	"strings"
	"testing"
)

func TestPingHost_Reachable(t *testing.T) {
	result := PingHost("google.com") 
	if strings.Contains(result, "unreachable") {
		t.Errorf("not expected 'unreachable' in result, got: %s", result)
	}
}

func TestPingHost_Unreachable(t *testing.T) {
	result := PingHost("invalid.invalid") 
	if !strings.Contains(result, "unreachable") {
		t.Errorf("expected 'unreachable' in result, got: %s", result)
	}
}
