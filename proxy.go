package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// createProxy creates a reverse proxy that forwards requests to Ollama
func createProxy(targetURL string) (*httputil.ReverseProxy, error) {
	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, fmt.Errorf("invalid Ollama URL: %w", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Override Director to modify the request
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = target.Host
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme
	}

	// Override ModifyResponse to log response details
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Get start time from request context
		var elapsed time.Duration
		if startTime, ok := resp.Request.Context().Value(requestStartTimeKey).(time.Time); ok {
			elapsed = time.Since(startTime)
		}

		// Log response status with elapsed time
		if elapsed > 0 {
			fmt.Printf("%s%sResponse: %s%s %s%s%s %s%s(%s)%s\n",
				colorGray, colorBold, colorReset,
				colorGreen, colorBold, resp.Status, colorReset,
				colorGray, colorReset, elapsed.Round(time.Millisecond), colorReset)
		} else {
			fmt.Printf("%s%sResponse: %s%s %s%s%s\n",
				colorGray, colorBold, colorReset,
				colorGreen, colorBold, resp.Status, colorReset)
		}
		fmt.Println()
		return nil
	}

	return proxy, nil
}
