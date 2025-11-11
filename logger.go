package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorMagenta = "\033[35m"
	colorGray   = "\033[90m"
)

// loggingMiddleware wraps an HTTP handler with beautiful request logging
func loggingMiddleware(next http.Handler, verbose bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Read request body if present
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
			r.Body.Close()
			// Restore body for the proxy
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Log request
		logRequest(r, bodyBytes, start, verbose)

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, body []byte, start time.Time, verbose bool) {
	timestamp := start.Format("15:04:05.000")
	
	if verbose {
		// Verbose mode: log full traffic
		fmt.Printf("\n%s%s╔════════════════════════════════════════════════════════════════╗%s\n", colorGray, colorBold, colorReset)
		fmt.Printf("%s%s║%s %s%sREQUEST%s %s%s│%s %s%s%s%s%s\n", colorGray, colorBold, colorReset, colorCyan, colorBold, colorReset, colorGray, colorBold, colorReset, colorGray, colorBold, timestamp, colorReset, colorReset)
		fmt.Printf("%s%s╠════════════════════════════════════════════════════════════════╣%s\n", colorGray, colorBold, colorReset)
		
		// Method
		fmt.Printf("%s%s║%s %sMethod:%s     %s%s%s%s\n", colorGray, colorBold, colorReset, colorYellow, colorReset, colorGreen, colorBold, r.Method, colorReset)
		
		// Path
		fmt.Printf("%s%s║%s %sPath:%s       %s%s%s%s\n", colorGray, colorBold, colorReset, colorYellow, colorReset, colorBlue, colorBold, r.URL.Path, colorReset)
		
		// Query string
		if r.URL.RawQuery != "" {
			fmt.Printf("%s%s║%s %sQuery:%s      %s%s%s\n", colorGray, colorBold, colorReset, colorYellow, colorReset, colorCyan, r.URL.RawQuery, colorReset)
		}
		
		// Headers
		if len(r.Header) > 0 {
			fmt.Printf("%s%s║%s %sHeaders:%s    %s\n", colorGray, colorBold, colorReset, colorYellow, colorReset, colorReset)
			for name, values := range r.Header {
				for _, value := range values {
					fmt.Printf("%s%s║%s            %s%s:%s %s%s%s\n", colorGray, colorBold, colorReset, colorMagenta, name, colorReset, colorCyan, value, colorReset)
				}
			}
		}
		
		// Body
		if len(body) > 0 {
			fmt.Printf("%s%s║%s %sBody:%s       %s\n", colorGray, colorBold, colorReset, colorYellow, colorReset, colorReset)
			
			// Try to pretty-print JSON
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "            ", "  "); err == nil {
				// Successfully formatted as JSON
				bodyLines := bytes.Split(prettyJSON.Bytes(), []byte("\n"))
				for _, line := range bodyLines {
					if len(line) > 0 {
						fmt.Printf("%s%s║%s%s%s\n", colorGray, colorBold, colorReset, string(line), colorReset)
					}
				}
			} else {
				// Not JSON, print as-is with line breaks
				bodyStr := string(body)
				if len(bodyStr) > 500 {
					bodyStr = bodyStr[:500] + "\n            ... (truncated)"
				}
				fmt.Printf("%s%s║%s            %s%s%s\n", colorGray, colorBold, colorReset, colorCyan, bodyStr, colorReset)
			}
		}
		
		fmt.Printf("%s%s╚════════════════════════════════════════════════════════════════╝%s\n\n", colorGray, colorBold, colorReset)
	} else {
		// Non-verbose mode: log only prompt content
		if len(body) > 0 {
			// Try to extract prompt from JSON body
			var jsonBody map[string]interface{}
			if err := json.Unmarshal(body, &jsonBody); err == nil {
				// Successfully parsed JSON
				if prompt, ok := jsonBody["prompt"].(string); ok && prompt != "" {
					fmt.Printf("\n%s%s[%s]%s %sPrompt:%s %s%s%s\n\n", 
						colorGray, colorBold, timestamp, colorReset,
						colorYellow, colorReset, colorCyan, prompt, colorReset)
					return
				}
			}
			// If no prompt found or not JSON, log nothing in non-verbose mode
		}
	}
}

