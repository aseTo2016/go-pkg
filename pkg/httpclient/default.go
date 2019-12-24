package httpclient

import "net/http"

// NewDefaultClient returns a new default http client
func NewDefaultClient() *Cli {
	c := new(Cli)
	c.nativeClient = http.DefaultClient
	return c
}
