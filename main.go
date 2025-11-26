package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getPrivateIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil && ip4.IsPrivate() {
				return ip4.String()
			}
		}
	}
	return ""
}

func main() {
	port := getEnv("OLLAMA_PROXY_PORT", "3131")
	// Ensure port starts with :
	if len(port) > 0 && port[0] != ':' {
		port = ":" + port
	}

	ollamaURL := getEnv("OLLAMA_HOST", "http://localhost:11434")

	// Create reverse proxy
	proxy, err := createProxy(ollamaURL)
	if err != nil {
		log.Fatalf("Failed to create proxy: %v", err)
	}

	// Create handler that wraps proxy with logging middleware
	handler := loggingMiddleware(proxy)

	// Get private IP address
	privateIP := getPrivateIP()

	// Start server
	msg := fmt.Sprintf("%s%sListening on: %shttp://localhost%s%s", colorGray, colorBold, colorBlue, port, colorReset)
	if privateIP != "" {
		msg += fmt.Sprintf(" & %shttp://%s%s%s", colorBlue, privateIP, port, colorReset)
	}
	msg += fmt.Sprintf("\n%sProxying to: %s%s%s\n", colorGray, colorCyan, ollamaURL, colorReset)
	fmt.Println(msg)

	// Bind to all interfaces (0.0.0.0) to allow remote connections
	if err := http.ListenAndServe("0.0.0.0"+port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
