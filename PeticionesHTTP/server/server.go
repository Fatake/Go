package server

import (
	"net/http"
)

type Server struct {
	puerto string
}

// NewServer
func NewServer(puerto string) *Server {
	return &Server{
		puerto: puerto,
	}
}

func (s *Server) Listen() error {
	err := http.ListenAndServe(s.puerto, nil)
	return err
}
