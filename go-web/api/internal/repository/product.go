package repository

import (
	"api/internal/model"
	"encoding/json"
	"os"
)

var products []model.Product

func LoadProducts(filePath string) error {
	byteValueJson, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValueJson, &products)
	if err != nil {
		return err
	}
	return nil

}

func GetProducts() []model.Product {
	return products
}

func AddProduct(newProduct model.Product, products []model.Product) {
	products = append(products, newProduct)
}
