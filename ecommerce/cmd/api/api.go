package api

import (
	"database/sql"
	"ecommerce/services/address"
	"ecommerce/services/cart"
	"ecommerce/services/order"
	"ecommerce/services/product"
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
	router := http.NewServeMux()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(router)

	addressStore := address.NewStore(s.db)
	addressHandler := address.NewHandler(addressStore)
	addressHandler.RegisterRoutes(router)

	orderStore := order.NewStore(s.db)
	orderHandler := order.NewHandler(orderStore, userStore)
	orderHandler.RegisterRoutes(router)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(router)

	log.Printf("Listening on %v", s.addr)

	return http.ListenAndServe(s.addr, router)
}
