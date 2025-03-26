package dto

type AddCartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
