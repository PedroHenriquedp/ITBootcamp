package service

import (
	"api/internal/model"
	"api/internal/repository"
	"fmt"
)

func UpdateProduct(updatedProduct model.Product) error {
	products := repository.GetProducts()

	for i, product := range products {
		if product.ID == updatedProduct.ID {
			products[i].Name = updatedProduct.Name
			products[i].Quantity = updatedProduct.Quantity
			products[i].CodeValue = updatedProduct.CodeValue
			products[i].IsPublished = updatedProduct.IsPublished
			products[i].Expiration = updatedProduct.Expiration
			products[i].Price = updatedProduct.Price

			return repository.SaveProducts(products)
		}
	}

	return fmt.Errorf("Produto com ID %d não encontrado", updatedProduct.ID)
}
