package repository

import (
	"encoding/json"
	"os"

	"github.com/PedroHenriquedp/api/internal"
)

type ProductRepository interface {
	FetchProducts() ([]internal.Product, error)
}

type JSONProductRepository struct {
	FilePath string
}

func (r *JSONProductRepository) FetchProducts() ([]internal.Product, error) {
	byteValueJson, err := os.ReadFile(r.FilePath)

	if err != nil {
		return nil, err
	}

	var products []internal.Product
	err = json.Unmarshal(byteValueJson, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
