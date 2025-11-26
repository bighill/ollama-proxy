# TODO

## Repo Health

- [ ] **CI/CD**: Add GitHub Actions workflow for running tests and linters on push/PR.
- [ ] **Linting**: Add `.golangci.yml` configuration for `golangci-lint`.
- [ ] **Makefile**: Create a `Makefile` with targets for `build`, `test`, `run`, and `lint`.
- [ ] **License**: Add a LICENSE file (e.g., MIT or Apache 2.0).

## Clean & Durable Code

- [ ] **Configuration**:
    - [ ] Make the listening port configurable via environment variable (e.g., `PORT` or `OLLAMA_PROXY_PORT`).
    - [ ] Make the upstream Ollama URL configurable via environment variable (e.g., `OLLAMA_HOST`).
- [ ] **Logging**:
    - [ ] Migrate from `fmt.Printf` to `log/slog` for structured logging (while maintaining the pretty console output for development).
    - [ ] Add log levels (INFO, DEBUG, ERROR).
- [ ] **Testing**:
    - [ ] Add unit tests for `logger.go` to verify JSON parsing and fallback behavior.
    - [ ] Add more comprehensive tests for `proxy.go` (e.g., testing response modification logic).
- [ ] **Error Handling**:
    - [ ] Improve error handling in `main.go` (graceful shutdown).
