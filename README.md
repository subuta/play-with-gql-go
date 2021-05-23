# play-with-gql-go
Playground for golang + grahql-go.

## How to develop

```bash
# Run main
$ go run .
```
## How to develop "WASM"

```bash
# Generate WASM version of main.
$ make build-wasm

# Run HTTP Server for WASM
$ go run ./cmds/http.go

# open WASM demo
open http://localhost:8080
```