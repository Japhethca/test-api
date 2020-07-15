package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	t.Run("getProducts", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		getProducts(res, req)

		if res.Result().StatusCode != http.StatusOK {
			t.Error("failed to fetch products")
		}

		var resData []Product
		json.NewDecoder(res.Body).Decode(&resData)

		if len(resData) != 1 {
			t.Error("invalid response")
		}
	})

	t.Run("sayHello", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		sayHello(res, req)

		if res.Result().StatusCode != http.StatusOK {
			t.Error("failed to fetch products")
		}

		var resMsg ResponseData
		json.NewDecoder(res.Body).Decode(&resMsg)

		if resMsg.Message != "Hello, this is a test server" {
			t.Error("invalid response")
		}
	})
}
