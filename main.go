package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":3500", nil)
}
