package handler

import (
	"net/http"

	"practice/internal/service"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router, service *service.OrderService) {
	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Заказы пока не реализованы"))
	}).Methods("GET")
}
