# rip

[![GoDoc](https://godoc.org/github.com/ripexz/rip?status.svg)](http://godoc.org/github.com/ripexz/rip)

Go package that can be used to get client's real public IP.
Based on [https://github.com/tomasen/realip](https://github.com/tomasen/realip).

### Features

- Parses IPs from `X-Real-IP`
- Parses IPs from `X-Forwarded-For`
- Excludes local/private address by default
- Custom filtering options for `X-Forwarded-For` IPs

## Examples

### Basic usage

```go
package main

import "github.com/ripexz/rip"

func (h *Handler) ServeIndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clientIP := rip.FromRequest(r, nil)
	log.Println("GET / from", clientIP)
}
```

### AWS ELB

```go
clientIP := rip.FromRequest(r, rip.FilterAWS)
```

### Custom Filtering

```go
clientIP := rip.FromRequest(r, func(ips []string) (string, bool) {
	// your custom logic here
	return "127.0.0.1", true
})
```

## Contributing

Please make sure your code:

- Passes the configured golangci-lint checks.
- Passes the tests.
