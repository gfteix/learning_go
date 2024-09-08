package main

import (
	"ecommerce/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
