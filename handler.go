package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	res := ResponseData{Message: "Hello, this is a test server"}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func jsonEncode(w io.Writer, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
