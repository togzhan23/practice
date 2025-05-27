package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"practice/internal/dal"
	"practice/internal/service"
	"practice/models"

	"github.com/gorilla/mux"
)

type SupplyService struct {
	repo dal.SupplyRepository
}

func NewSupplyService(repo dal.SupplyRepository) *SupplyService {
	return &SupplyService{repo: repo}
}

func (s *SupplyService) GetAll(ctx context.Context) ([]models.FarmSupplyItem, error) {
	supplies, err := s.repo.GetAll()
	if err != nil {
		log.Printf("ошибка получения поставок: %v", err)
		return nil, err
	}
	return supplies, nil
}

func RegisterSupplyRoutes(router *mux.Router, svc *service.SupplyService) {
	router.HandleFunc("/supplies", func(w http.ResponseWriter, r *http.Request) {
		supplies, err := svc.GetAll(r.Context())
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(supplies)
	}).Methods("GET")
}
