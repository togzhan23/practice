package service

import (
	"practice/internal/dal"
)

type OrderService struct {
	repo *dal.OrderRepository
}

func NewOrderService(repo *dal.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}
