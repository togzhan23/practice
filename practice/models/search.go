package models

type SearchFilter struct {
	Query      string
	Category   string
	PriceRange [2]float64
}
