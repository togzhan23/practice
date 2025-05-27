package models

type SupplyOrder struct {
	OrderID   int
	ProductID int
	Quantity  float64
	Status    string
}

type CropStat struct {
	CropName      string  `json:"crop_name"`
	TotalQuantity float64 `json:"total_quantity"`
}
