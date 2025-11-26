# Ollama Streaming Proxy

A lightweight, zero-dependency Go HTTP proxy service that forwards all requests to Ollama and provides beautiful, colorized logs of requests and responses.

## Features

- **Zero dependencies**: Uses only the Go standard library.
- **Transparent proxying**: Forwards all requests to Ollama with full streaming support.
- **Beautiful logging**:
  - Logs request timestamp and content (extracts prompt from JSON).
  - Logs response status and duration.
  - Color-coded output for better readability.
- **Remote access**: Binds to `0.0.0.0` to allow access from other devices on your local network.

## Usage

### Prerequisites

- [Go](https://go.dev/dl/) installed.
- [Ollama](https://ollama.ai/) running on `http://localhost:11434`.

### Running the Proxy

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ollama-proxy.git
   cd ollama-proxy
   ```

2. Run the proxy:
   ```bash
   go run .
   ```

3. The proxy will start on `http://localhost:3131`. It will also display your private IP address for remote access.

### Building

To build a binary for production or distribution:

```bash
go build -o ollama-proxy .
./ollama-proxy
```

### Using the Proxy

Send requests to the proxy just like you would to Ollama:

```bash
curl http://localhost:3131/api/generate -d '{"model": "llama2", "prompt": "Hello"}'
```

## Configuration

The proxy can be configured using environment variables:
- **`OLLAMA_PROXY_PORT`**: The port the proxy listens on (default: `3131`).
- **`OLLAMA_HOST`**: The URL of the upstream Ollama instance (default: `http://localhost:11434`).

Example:
```bash
OLLAMA_PROXY_PORT=8080 OLLAMA_HOST=http://192.168.1.100:11434 ./ollama-proxy
```

## Development

See [AGENTS.md](AGENTS.md) for development guidelines.

