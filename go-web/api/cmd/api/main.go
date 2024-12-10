package main

import (
	"api/internal/repository"
	"api/internal/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := repository.LoadProducts("products.json")
	if err != nil {
		log.Fatalf("Erro ao carregar produtos: %v", err)
	}

	rt := chi.NewRouter()
	rt.Mount("/", routes.ProductRoutes())

	log.Fatal(http.ListenAndServe(":8080", rt))
}
