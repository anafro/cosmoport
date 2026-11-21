package server

import (
	"net/http"

	"github.com/anafro/cosmoport/internal/config"
)

type Server struct {
	Config   *config.Config
	Handlers map[string]http.Handler
	Logs     chan string
}

func NewServer(conf *config.Config) *Server {
	server := &Server{
		Config:   conf,
		Handlers: make(map[string]http.Handler),
		Logs:     make(chan string),
	}

	server.SyncHandlers()
	return server
}

func (s *Server) SyncHandlers() {
	clear(s.Handlers)

	for _, rule := range s.Config.Rules {
		destination, err := rule.Destination()

		if err != nil {
			panic(err)
		}

		var handler http.Handler
		switch destination.Type {
		case "backend":
			handler = NewBackendReverseProxy(rule.Domain, rule.Backend)
		case "static":
			handler = NewStaticProxy(rule.Domain, rule.Static)
		default:
			panic("unknown destination type: " + destination.Type)
		}

		s.Handlers[rule.Domain] = handler
	}
}

func (s *Server) Start() {
	err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for domain, handler := range s.Handlers {
			if r.Host == domain {
				handler.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	}))

	if err != nil {
		s.Logs <- err.Error()
	}
}
