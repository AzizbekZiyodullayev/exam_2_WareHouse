package models

type RespProduct struct {
	BranchID   string  `json:"branch_id"`
	ComingID   string  `json:"coming_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}
