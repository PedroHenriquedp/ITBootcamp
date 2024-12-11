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

func AddProduct(newProduct model.Product) {
	products = append(products, newProduct)
}

func SaveProducts(products []model.Product) error {
	byteValueJson, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = os.WriteFile("/Users/peddpereira/Desktop/Project/bootcamp/go-web/api/cmd/api/products.json", byteValueJson, 0644)
	return err
}
