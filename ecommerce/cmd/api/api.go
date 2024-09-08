package api

import (
	"database/sql"
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
	return http.ListenAndServe(s.addr, mux)
}
