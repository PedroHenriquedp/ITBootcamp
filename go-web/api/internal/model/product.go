package model

type Product struct {
	ID          uint16  `json:"id"`
	Name        string  `json:"name"`
	Quantity    uint16  `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
