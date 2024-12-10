package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

// Criando estrutura para desserealizar o json
type NameRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`

}

func main() {
	rt := chi.NewRouter()

	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var nameReq NameRequest

		if err := json.NewDecoder(r.Body).Decode(&nameReq);

		err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Hello %s %s/n", nameReq.FirstName, nameReq.LastName)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))

	})

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
