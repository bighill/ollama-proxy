package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	colorReset = "\033[0m"
	colorBold  = "\033[1m"
	colorGreen = "\033[32m"
	colorBlue  = "\033[34m"
	colorCyan  = "\033[36m"
	colorGray  = "\033[90m"
)

type contextKey string

const requestStartTimeKey contextKey = "requestStartTime"

// loggingMiddleware wraps an HTTP handler with beautiful request logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Store start time in request context for response timing
		ctx := context.WithValue(r.Context(), requestStartTimeKey, start)
		r = r.WithContext(ctx)

		// Read request body if present
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
			r.Body.Close()
			// Restore body for the proxy
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Log request
		logRequest(bodyBytes, start)

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

func logRequest(body []byte, start time.Time) {
	timestamp := start.Format("15:04:05.000")

	if len(body) > 0 {
		var payload struct {
			Messages []struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"messages"`
		}
		if err := json.Unmarshal(body, &payload); err == nil && len(payload.Messages) > 0 {
			fmt.Printf("%s%s%s\n", colorGray, colorBold, timestamp)
			fmt.Printf("%s%s%s\n", colorBlue, colorBold, payload.Messages[0].Content)
		} else {
			fmt.Printf("%s%s%s\n", colorGray, colorBold, timestamp)
			fmt.Println("Content not available")
		}
	} else {
		fmt.Printf("%s%s%s\n", colorGray, colorBold, timestamp)
	}
}
