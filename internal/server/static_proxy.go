package server

import "net/http"

func NewStaticProxy(domain, path string) http.Handler {
	return http.FileServer(http.Dir(path))
}
