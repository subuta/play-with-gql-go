# play-with-gql-go
Playground for golang + grahql-go.

Includes these examples.

- [x] How to use "github.com/graphql-go/graphql"
- [x] How to use "github.com/mattn/go-sqlite3"
- [x] How to run generated WASM in Browser.
  -  `go-sqlite3` won't work on WASM build.

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