package routes

import (
	"api/internal/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProductRoutes() chi.Router {
	rt := chi.NewRouter()

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	rt.Get("/products", controller.GetProductsHandler)
	rt.Get("/products/{id}", controller.GetProductsByIDHandler)
	rt.Post("/products", controller.CreateProductHandler)

	return rt
}
