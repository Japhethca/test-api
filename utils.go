package main

import (
	"encoding/json"
	"net/http"
)

func jsonEncode(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
