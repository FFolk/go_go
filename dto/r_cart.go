package dto

import "time"

type R_Cart struct {
	CartID     int                   `json:"cart_id"`
	CustomerID int                   `json:"customer_id"`
	CartName   string                `json:"cart_name"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
	CartItem   []CartItemWithProduct `json:"cart_item"`
}

type CartItemWithProduct struct {
	CartItemID  int     `json:"cart_item_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       string  `json:"price"`
	Quantity    int     `json:"quantity"`
	ItemTotal   float64 `json:"item_total"`
}
