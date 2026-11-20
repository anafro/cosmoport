package server

import "github.com/anafro/cosmoport/internal/config"

type Server struct {
	Config *config.Config
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		Config: conf,
	}
}

func (s *Server) Start() {

}
