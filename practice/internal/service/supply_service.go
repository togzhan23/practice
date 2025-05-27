package service

import (
	"context"
	"practice/internal/dal"
	"practice/models"
)

type SupplyService struct {
	repo *dal.SupplyRepository
}

func NewSupplyService(repo *dal.SupplyRepository) *SupplyService {
	return &SupplyService{repo: repo}
}

func (s *SupplyService) ListSupplyItems() ([]models.FarmSupplyItem, error) {
	return s.repo.GetAll()
}
func (s *SupplyService) GetAll(ctx context.Context) ([]models.FarmSupplyItem, error) {
	supplies, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return supplies, nil
}
