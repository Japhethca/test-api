package main

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ID          int    `json:"id"`
}

var products = []Product{
	{
		ID:          1,
		Name:        "Galaxy s20",
		Description: "A smartphone",
		Price:       20,
	},
}
