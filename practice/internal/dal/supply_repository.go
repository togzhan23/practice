package dal

import (
	"database/sql"
	"practice/models"
)

type SupplyRepository struct {
	db *sql.DB
}

func NewSupplyRepository(db *sql.DB) *SupplyRepository {
	return &SupplyRepository{db: db}
}

func (r *SupplyRepository) GetAll() ([]models.FarmSupplyItem, error) {
	return []models.FarmSupplyItem{}, nil
}
