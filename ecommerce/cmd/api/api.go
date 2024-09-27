package api

import (
	"database/sql"
	"ecommerce/services/user"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// https://codewithflash.com/advanced-routing-with-go-122

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	// how to create a subrouter for /users?
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(mux)

	log.Printf("Listening on %v", s.addr)

	return http.ListenAndServe(s.addr, mux)

}
