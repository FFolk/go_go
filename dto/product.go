package dto

import "time"

type Product struct {
	ProductID     int       `json:"product_id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	Price         string    `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
