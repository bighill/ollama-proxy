package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const ollamaURL = "http://localhost:11434"

// createProxy creates a reverse proxy that forwards requests to Ollama
func createProxy() (*httputil.ReverseProxy, error) {
	target, err := url.Parse(ollamaURL)
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
		// Log response status
		fmt.Printf("%s%sResponse: %s%s %s%s%s\n\n", 
			colorGray, colorBold, colorReset,
			colorGreen, colorBold, resp.Status, colorReset)
		return nil
	}

	return proxy, nil
}

