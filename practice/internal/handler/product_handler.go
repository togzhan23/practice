package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"practice/internal/service"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router, service *service.ProductService) {
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := service.GetAll(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	}).Methods("GET")
}
