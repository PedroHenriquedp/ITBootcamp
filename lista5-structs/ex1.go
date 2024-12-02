package main

import "fmt"

type Product struct {
    ID          int
    Name        string
    Price       float64
    Description string
    Category    string
}

var Products []Product

func (p *Product) Save() {
    Products = append(Products, *p)
}

func GetAll() {
    for _, product := range Products {
        fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
            product.ID, product.Name, product.Price, product.Description, product.Category)
    }
}

func getById(id int) *Product {
    for _, product := range Products {
        if product.ID == id {
            return &product
        }
    }
    return nil
}

func main() {
    product1 := Product{ID: 1, Name: "Chocolate", Price: 10.0, Description: "Delicious dark chocolate", Category: "Candy"}
    product2 := Product{ID: 2, Name: "Soda", Price: 5.0, Description: "Refreshing soda", Category: "Beverage"}
    
    product1.Save()
    product2.Save()

    fmt.Println("Produtos disponíveis:")
    GetAll()

    idToFind := 1
    foundProduct := getById(idToFind)
    if foundProduct != nil {
        fmt.Printf("Produto encontrado: ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
            foundProduct.ID, foundProduct.Name, foundProduct.Price, foundProduct.Description, foundProduct.Category)
    } else {
        fmt.Printf("Produto com ID %d não encontrado.\n", idToFind)
    }
}

