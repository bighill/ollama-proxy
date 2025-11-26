# Agent Guide

This repository is a lightweight, zero-dependency Go HTTP proxy for Ollama. It is designed to be simple, transparent, and easy to maintain.

## Philosophy

- **Zero Dependencies**: We strictly avoid external dependencies. The standard library is sufficient for our needs.
- **Simplicity**: Code should be easy to read and understand. Avoid over-engineering.
- **Transparency**: The proxy should not modify the request/response body unless necessary for logging.

## Architecture

- **`main.go`**: Entry point. Sets up the server, handles signal interrupts (if added), and starts the listener.
- **`proxy.go`**: Contains the reverse proxy logic (`httputil.ReverseProxy`). It handles the `Director` (request modification) and `ModifyResponse` (response logging).
- **`logger.go`**: Handles request logging, including body parsing (to extract prompts) and colorized output.

## Development

### Testing

- **Unit Tests**: Run `go test ./...` to run all unit tests.
- **Integration Tests**: Use `./test_proxy.sh` to verify the proxy against a running Ollama instance.

### Code Style

- Follow standard Go formatting (`gofmt`).
- Ensure all exported functions and types have comments.
- Keep functions small and focused.

## Common Tasks

- **Adding Features**: If adding a feature requires a new flag, consider if it's strictly necessary. We prefer convention over configuration for this tool.
- **Debugging**: The proxy logs to stdout. Use `fmt.Printf` for debugging if needed, but remove it before committing.
