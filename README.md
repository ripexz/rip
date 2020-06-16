# rip

[![GoDoc](https://godoc.org/github.com/ripexz/rip?status.svg)](http://godoc.org/github.com/ripexz/rip)

Go package that can be used to get client's real public IP.
Based on [https://github.com/tomasen/realip](https://github.com/tomasen/realip).

### Features

- Follows the rule of X-Real-IP
- Follows the rule of X-Forwarded-For
- Exclude local or private address

## Example

```go
package main

import "github.com/ripexz/rip"

func (h *Handler) ServeIndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clientIP := rip.FromRequest(r, nil)
	log.Println("GET / from", clientIP)
}
```

## Contributing

Please make sure your code:

- Passes the configured golangci-lint checks.
- Passes the tests.
