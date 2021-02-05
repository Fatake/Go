package main

import "net/http"

// Router Estructura
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// NewRouter Crea un router nuevo
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// FindHandler Encuentra el hanled
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

// FindPath Encuentra el Path
func (r *Router) FindPath(path string) bool {
	_, exist := r.rules[path]
	return exist
}

// ServeHTTP Pone a escucha
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
