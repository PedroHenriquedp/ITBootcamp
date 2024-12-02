package main

import "fmt"

type Product interface {
    Price() float64
}

type Small struct {
    price float64
}

type Medium struct {
    price float64
}

type Large struct {
    price float64
}

func (s Small) Price() float64 {
    return s.price
}

func (m Medium) Price() float64 {
    return m.price + (m.price * 0.03) + (m.price * 0.03)
}

func (l Large) Price() float64 {
    return l.price + (l.price * 0.06) + 2500
}

func CreateProduct(productType string, price float64) Product {
    switch productType {
    case "small":
        return Small{price: price}
    case "medium":
        return Medium{price: price}
    case "large":
        return Large{price: price}
    default:
        return nil
    }
}

func main() {
    smallProduct := CreateProduct("small", 100)
    mediumProduct := CreateProduct("medium", 200)
    largeProduct := CreateProduct("large", 300)

    fmt.Printf("Preço do produto pequeno: %.2f\n", smallProduct.Price())
    fmt.Printf("Preço do produto médio: %.2f\n", mediumProduct.Price())
    fmt.Printf("Preço do produto grande: %.2f\n", largeProduct.Price())
}

