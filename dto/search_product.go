package dto

type SearchProduct struct {
	Pricemin    string `json:"price_min"`
	Pricemax    string `json:"price_max"`
	Description string `json:"description"`
}
