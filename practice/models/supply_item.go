package models

type FarmSupplyItem struct {
	SupplyID    int                    `json:"supply_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Unit        string                 `json:"unit"`
	Quantity    float64                `json:"quantity"`
	Category    string                 `json:"category"`
	Metadata    map[string]interface{} `json:"metadata"`
}
type Supply struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
