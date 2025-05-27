package main

import (
	"database/sql"
	"log"
	"net/http"
	"practice/internal/dal"
	"practice/internal/handler"
	"practice/internal/service"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=agriculture sslmode=disable")
	if err != nil {
		log.Fatal("DB error:", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	productRepo := dal.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	handler.RegisterProductRoutes(router, productService)

	supplyRepo := dal.NewSupplyRepository(db)
	supplyService := service.NewSupplyService(supplyRepo)

	handler.RegisterSupplyRoutes(router, supplyService)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	orderRepo := dal.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	handler.RegisterOrderRoutes(router, orderService)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AgriPlatform API работает "))
	})

	log.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", router)
}
