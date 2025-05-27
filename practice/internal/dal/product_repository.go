package dal

import (
	"database/sql"
	"practice/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	// реализация запроса к БД
	return []models.Product{}, nil
}
