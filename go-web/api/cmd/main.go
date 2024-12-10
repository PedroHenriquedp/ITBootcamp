package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          uint16  `json:"id"`
	Name        string  `json:"name"`
	Quantity    uint16  `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func isValidDate(date string) bool {
	_, err := time.Parse("02/01/2006", date)
	return err == nil
}

func main() {
	byteValueJson, err := os.ReadFile("/Users/peddpereira/Desktop/Project/bootcamp/go-web/lista2-interagindoAPI/docs/db/products.json")
	if err != nil {
		fmt.Print("Erro ao ler o arquivo %v\n", err)
		return
	}
	fmt.Printf("Sucesso ao ler o arquivo!!\n")

	var products []Product
	err = json.Unmarshal(byteValueJson, &products)
	if err != nil {
		fmt.Printf("Deu chabu! O erro foi:%v\n", err)
		return
	}

	for _, product := range products {
		fmt.Printf("id = %d, name = %s, quantidade = %d\n", product.ID, product.Name, product.Quantity)
	}

	rt := chi.NewRouter()

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	rt.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		for _, product := range products {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(product.Name))
			for _, product := range products {
				json.NewEncoder(w).Encode(product)
			}
		}
	})

	rt.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		idProduct, err := strconv.ParseUint(idParam, 10, 16)
		if err != nil {
			http.Error(w, "ID invalido", http.StatusBadRequest)
			return
		}

		for _, product := range products {
			if product.ID == uint16(idProduct) {
				w.Write([]byte(fmt.Sprintf("O produto é \n%s", product.Name)))
				return
			}
		}
	})

	rt.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		var newProduct Product

		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Produto inválido", http.StatusBadRequest)
			return
		}

		if newProduct.Name == "" || newProduct.CodeValue == "" || newProduct.Expiration == "" || newProduct.Price <= 0 {
			http.Error(w, "Name, CodeValue, Expiration e Price precisam estar preenchidos", http.StatusBadRequest)
			return
		}

		for _, product := range products {
			if product.CodeValue == newProduct.CodeValue {
				http.Error(w, "O produto precisa ter um CodeValue único", http.StatusBadRequest)
				return
			}
		}

		if !isValidDate(newProduct.Expiration) {
			http.Error(w, "Data inválida", http.StatusBadRequest)
			return
		}

		var maxID uint16 = 0
		for _, product := range products {
			if product.ID > maxID {
				maxID = product.ID
			}
		}
		newProduct.ID = maxID + 1

		products = append(products, newProduct)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newProduct)
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
