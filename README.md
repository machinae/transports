# transports
Simple, idiomatic HTTP Client Middleware for Go

Transports provides an easy way to wrap HTTP clients with chainable middleware
for timeouts, logging, proxies, and other common functionality.

## Install
`go get github.com/machinae/transports`

## Quick Start
Just wrap any existing `http.RoundTripper` in the Transport struct you want to
use.

```go
import (
  "fmt"
  "net/http"
  "github.com/machinae/transports"
)

  func main() {
    tr := &transports.LoggingTransport{http.DefaultTransport}

    c := &http.Client{
      Transport: tr,
    }

    resp, err := c.Get("http://www.example.com")
    fmt.Println(resp, err)
  }
)
```

## Configuration
The zero value of any Transport in this package is a valid Transport. To
configure a Transport, just set the struct fields before setting it on the
client.

```go
var tr http.RoundTripper
tr = &transports.TimeoutTransport{Timeout: 15 * time.Second}
// Transports can be chained like middleware at any time
tr = &transports.LoggingTransport{tr}
```
