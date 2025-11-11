package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":3131"

func main() {
	// Create reverse proxy
	proxy, err := createProxy()
	if err != nil {
		log.Fatalf("Failed to create proxy: %v", err)
	}

	// Create handler that wraps proxy with logging middleware
	handler := loggingMiddleware(proxy)

	// Start server
	fmt.Printf("%s%s╔════════════════════════════════════════════════════════════════╗%s\n", colorGray, colorBold, colorReset)
	fmt.Printf("%s%s║%s %s%sOllama Proxy Server%s                                    %s%s║%s\n", colorGray, colorBold, colorReset, colorCyan, colorBold, colorReset, colorGray, colorBold, colorReset)
	fmt.Printf("%s%s╠════════════════════════════════════════════════════════════════╣%s\n", colorGray, colorBold, colorReset)
	fmt.Printf("%s%s║%s Listening on: %s%shttp://localhost%s%s%s%s\n", colorGray, colorBold, colorReset, colorBlue, colorBold, port, colorReset, colorGray, colorBold, colorReset)
	fmt.Printf("%s%s║%s Proxying to:   %s%s%s%s%s%s\n", colorGray, colorBold, colorReset, colorGreen, colorBold, ollamaURL, colorReset, colorGray, colorBold, colorReset)
	fmt.Printf("%s%s╚════════════════════════════════════════════════════════════════╝%s\n\n", colorGray, colorBold, colorReset)

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

