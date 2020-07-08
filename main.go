package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/products", getProducts)
	http.ListenAndServe(":3500", nil)
}
