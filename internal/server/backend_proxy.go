package server

import (
	"net/http/httputil"

	"github.com/anafro/cosmoport/internal/support"
)

func NewBackendReverseProxy(domain string, sbackend string) *httputil.ReverseProxy {
	backend := support.Panics(ParseBackend(sbackend))
	url := support.Panics(backend.ToURL())
	return httputil.NewSingleHostReverseProxy(url)
}
