package dto

import "time"

type CartItem struct {
	CartItemID int       `json:"cart_item_id"`
	CartID     int       `json:"cart_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Cart       Cart      `json:"cart"`
}
