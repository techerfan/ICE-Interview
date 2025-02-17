package entity

type CartItem struct {
	ID          uint    `json:"id"`
	CartID      uint    `json:"cart_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
