package main

import (
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	res := ResponseData{Message: "Hello, this is a test server"}

	w.WriteHeader(http.StatusOK)
	jsonEncode(w, &res)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonEncode(w, &products)
}
