package models

import "github.com/google/uuid"

type Cart struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	SessionID uuid.UUID  `db:"session_id" json:"session_id"`
	CartItems []CartItem `json:"cart_items"`
}

type CartItem struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CartID    uuid.UUID `db:"cart_id" json:"cart_id"`
	ProductID string    `db:"product_id" json:"product_id"`
	Product   Product   `json:"product"`
	Quantity  int       `db:"quantity" json:"quantity"`
}
