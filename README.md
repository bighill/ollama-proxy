# Ollama Streaming Proxy

A lightweight Go HTTP proxy service that forwards all requests to Ollama and logs them with beautiful, readable terminal output.

## Features

- **Zero dependencies** - Uses only Go standard library
- **Transparent proxying** - Forwards all requests to Ollama with streaming support
- **Beautiful logging** - Color-coded, formatted request logging to stdout
- **Streaming support** - Automatically handles chunked encoding for streaming responses

## Usage

1. Make sure Ollama is running on `http://localhost:11434`

2. Run the proxy:
   ```bash
   go run .
   ```

3. The proxy will start on `http://localhost:3131`

4. Send requests to the proxy instead of directly to Ollama:
   ```bash
   curl http://localhost:3131/api/generate -d '{"model": "llama2", "prompt": "Hello"}'
   ```

## Configuration

Currently hardcoded:
- Proxy port: `3131`
- Ollama URL: `http://localhost:11434`

## Project Structure

- `main.go` - HTTP server setup and entry point
- `proxy.go` - Reverse proxy configuration
- `logger.go` - Pretty request logging middleware
- `go.mod` - Go module definition

