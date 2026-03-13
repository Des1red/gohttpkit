## gohttpkit

Minimal HTTP runtime utilities for Go.

### Provides:

- runtime environment loader

- CORS middleware

- middleware chaining

- simple handler composition

- No framework. Works with net/http.

### Installation
```go
go get github.com/Des1red/gohttpkit
```
### Environment

Create a .env file.
```sh
DEV=true
PORT=8080

DEV_ORIGIN=http://localhost:5173
PROD_ORIGIN=https://example.com
```
#### Required variables:
```
Variable	Description
DEV	        true or false
PORT	    server port
API_DEV	    dev API URL
API_PROD	production API URL
DEV_ORIGIN	allowed dev origin
PROD_ORIGIN	allowed production origin
```

### Basic Usage
```go
package main

import (
	"net/http"

	"github.com/Des1red/gohttpkit/httpkit"
)

func main() {

	runtime := httpkit.BuildRuntime(".env")

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	handler := httpkit.With(runtime, mux)

	http.ListenAndServe(runtime.Port, handler)
}
```

#### Ignore CORS for Paths
```go
httpkit.CORSIgnorePaths(
	"/auth",
	"/health",
)
```
Paths are matched using prefix rules.

Example:

/auth
/auth/login
/auth/refresh

#### Configure Allowed Methods

Default:

- GET, OPTIONS

Custom methods:
```go
httpkit.CORSMethods("GET, POST, PUT, DELETE, OPTIONS")
```

Allow all:
```go
httpkit.CORSMethods("")
Middleware Chaining
handler := httpkit.Chain(
	mux,
	httpkit.CORS(runtime),
)
```
Equivalent to:
```go
handler := httpkit.CORS(runtime)(mux)
Runtime Structure
type Runtime struct {
	Dev        bool
	DevOrigin  string
	ProdOrigin string
	Port       string
}
```

Created with:
```go
runtime := httpkit.BuildRuntime(".env")
```
### Philosophy

- minimal surface

- explicit behavior

- zero framework magic

- built on net/http

### Author

Des1red

### LISCENSE 

 MIT 