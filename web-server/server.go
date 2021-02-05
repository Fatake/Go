package main

import (
	"fmt"
	"log"
	"net/http"
)

// Server Structura del servidor
type Server struct {
	port   string
	router *Router
}

// NewServer Crea un nuevo Servidor
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Listen Pone a escuchar el servidor
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		log.Fatal("[!] Inposible de lanzar el servidor", err)
		fmt.Println("[!] Inposible de lanzar el servidor", err)
		return err
	}
	fmt.Println("[+] Server Activo", s.port)
	log.Print("[+] Server Activo", s.port)
	return nil
}

// Handle Pone a escuchar
func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	//Check if the path already exists
	if !s.router.FindPath(path) {
		//If not path then create a new one
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// AddMiddleware AddMiddleware
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Handler Handler
type Handler func(w http.ResponseWriter, r *http.Request)

// Middleware Middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc
