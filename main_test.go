package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateProxy(t *testing.T) {
	proxy, err := createProxy("http://localhost:11434")
	if err != nil {
		t.Fatalf("createProxy() failed: %v", err)
	}
	if proxy == nil {
		t.Fatal("createProxy() returned nil proxy")
	}
}

func TestGetPrivateIP(t *testing.T) {
	ip := getPrivateIP()
	// Just verify it doesn't crash and returns something reasonable
	// It might return empty string if no private IP is found, which is OK
	if ip != "" {
		t.Logf("Found private IP: %s", ip)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	// Create a simple handler that echoes the request body
	echoHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	// Wrap with logging middleware
	wrapped := loggingMiddleware(echoHandler)

	// Create test request with body
	body := `{"test": "data"}`
	req := httptest.NewRequest("POST", "/api/test", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Serve request
	wrapped.ServeHTTP(rr, req)

	// Verify response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Verify body was preserved
	if rr.Body.String() != body {
		t.Errorf("Expected body %q, got %q", body, rr.Body.String())
	}
}

func TestLoggingMiddlewarePreservesBody(t *testing.T) {
	var receivedBody []byte

	// Handler that captures the body
	captureHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	})

	wrapped := loggingMiddleware(captureHandler)

	body := `{"model": "llama2", "prompt": "test"}`
	req := httptest.NewRequest("POST", "/api/generate", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if string(receivedBody) != body {
		t.Errorf("Body not preserved. Expected %q, got %q", body, string(receivedBody))
	}
}

func TestLoggingMiddlewareWithEmptyBody(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrapped := loggingMiddleware(handler)

	req := httptest.NewRequest("GET", "/api/test", nil)
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}

func TestLogRequest(t *testing.T) {
	// Just verify logRequest doesn't crash with various inputs
	body := `{"test": "json"}`
	req := httptest.NewRequest("POST", "/api/test?foo=bar", bytes.NewBufferString(body))
	req.Header.Set("X-Test-Header", "test-value")

	start := time.Now()
	logRequest([]byte(body), start)
	// If we get here without panic, test passes
}
