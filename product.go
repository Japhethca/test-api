package main

import "errors"

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Price       int      `json:"price"`
	ID          int      `json:"id"`
}

var (
	errProductExists = errors.New("Product already exists")
)

var products = []Product{
	{
		ID:          1,
		Name:        "Galaxy s20",
		Description: "A smartphone",
		Categories:  []string{"smartphone"},
		Price:       20,
	},
}

// func createProduct(product Product) (error, *Product) {
// 	for _, p := range products {
// 		if p.Name == product.Name {
// 			return errProductExists, &Product{}
// 		}
// 	}
// 	return
// }
