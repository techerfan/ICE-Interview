package entity

type CartStatus string

const (
	CartOpen   CartStatus = "open"
	CartClosed CartStatus = "closed"
)

type CartEntity struct {
	ID        uint       `json:"id"`
	Total     float64    `json:"total"`
	SessionID string     `json:"session_id"`
	Status    CartStatus `json:"status"`
}
