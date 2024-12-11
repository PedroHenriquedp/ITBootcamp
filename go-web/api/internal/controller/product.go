package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"api/internal/model"
	"api/internal/repository"
	"api/internal/service"

	"github.com/go-chi/chi/v5"
)

func isValidDate(date string) bool {
	_, err := time.Parse("02/01/2006", date) //quero mudar aqui
	return err == nil
}

func findProductByID(id uint16, products []model.Product) *model.Product { //quero mudar aqui
	for i := range products {
		if products[i].ID == id {
			return &products[i]
		}
	}
	return nil
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := repository.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductsByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idProduct, err := strconv.ParseUint(idParam, 10, 16)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, product := range repository.GetProducts() {
		if product.ID == uint16(idProduct) {
			w.Write([]byte(fmt.Sprintf("O produto é \n%s", product.Name)))
			return
		}
	}
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product

	products := repository.GetProducts()

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
	}

	var maxID uint16 = 0
	for _, product := range repository.GetProducts() {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	newProduct.ID = maxID + 1

	repository.AddProduct(newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Produto inválido", http.StatusBadRequest)
		return
	}

	products := repository.GetProducts()
	existProduct := findProductByID(newProduct.ID, products)

	if existProduct == nil {
		http.Error(w, "Produto inexistente", http.StatusNotFound)
	}

	if newProduct.Name == "" || newProduct.CodeValue == "" || newProduct.Expiration == "" || newProduct.Price <= 0 {
		http.Error(w, "Name, CodeValue, Expiration e Price precisam estar preenchidos", http.StatusBadRequest)
		return
	}

	if !isValidDate(newProduct.Expiration) {
		http.Error(w, "Data inválida", http.StatusBadRequest)
	}

	err := service.UpdateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("200 Deu certo!!"))

}
