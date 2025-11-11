package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

const port = ":3131"

// getPrivateIP returns the first private IP address found on the system
func getPrivateIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		// Skip loopback and down interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Check if it's a private IP (RFC 1918)
			if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
				if ip.IsPrivate() {
					return ip.String()
				}
			}
		}
	}

	return ""
}

func main() {
	// Parse command line flags
	verbose := flag.Bool("v", false, "Enable verbose logging (show full traffic)")
	flag.Parse()

	// Create reverse proxy
	proxy, err := createProxy()
	if err != nil {
		log.Fatalf("Failed to create proxy: %v", err)
	}

	// Create handler that wraps proxy with logging middleware
	handler := loggingMiddleware(proxy, *verbose)

	// Get private IP address
	privateIP := getPrivateIP()

	// Start server
	fmt.Printf("%s%s╔════════════════════════════════════════════════════════════════╗%s\n", colorGray, colorBold, colorReset)
	fmt.Printf("%s%s║%s %s%sOllama Proxy Server%s                                    %s%s║%s\n", colorGray, colorBold, colorReset, colorCyan, colorBold, colorReset, colorGray, colorBold, colorReset)
	fmt.Printf("%s%s╠════════════════════════════════════════════════════════════════╣%s\n", colorGray, colorBold, colorReset)
	fmt.Printf("%s%s║%s Listening on: %s%shttp://localhost%s%s%s%s║%s\n", colorGray, colorBold, colorReset, colorBlue, colorBold, port, colorReset, colorGray, colorBold, colorReset)
	if privateIP != "" {
		fmt.Printf("%s%s║%s              %s%shttp://%s%s%s%s%s║%s\n", colorGray, colorBold, colorReset, colorBlue, colorBold, privateIP, port, colorReset, colorGray, colorBold, colorReset)
	}
	fmt.Printf("%s%s║%s Proxying to:   %s%s%s%s%s%s║%s\n", colorGray, colorBold, colorReset, colorGreen, colorBold, ollamaURL, colorReset, colorGray, colorBold, colorReset)
	fmt.Printf("%s%s╚════════════════════════════════════════════════════════════════╝%s\n\n", colorGray, colorBold, colorReset)

	// Bind to all interfaces (0.0.0.0) to allow remote connections
	if err := http.ListenAndServe("0.0.0.0"+port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
