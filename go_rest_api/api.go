package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// Initialize the router and register its dependencies
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering our services
	userService := NewUserService(s.store)
	userService.RegisterRoutes(subrouter)

	projectService := NewProjectService(s.store)
	projectService.RegisterRoutes(subrouter)

	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(subrouter)

	log.Println("Starting the API Server at ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
