package service

import (
	"context"
	"log"
	"practice/internal/dal"
	"practice/models"
)

type ProductService struct {
	repo dal.ProductRepository
}

func NewProductService(repo dal.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		log.Printf("ошибка получения продуктов: %v", err)
		return nil, err
	}
	return products, nil
}
