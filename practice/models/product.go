package models

type CropProduct struct {
	ProductID      int                    `json:"product_id"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	PricePerUnit   float64                `json:"price_per_unit"`
	Categories     []string               `json:"categories"`
	Certifications []string               `json:"certifications"`
	Components     []CropComponent        `json:"components"`
	Metadata       map[string]interface{} `json:"metadata"`
}

type CropComponent struct {
	ProductID   int     `json:"product_id"`
	ComponentID int     `json:"component_id"`
	Quantity    float64 `json:"quantity"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}
