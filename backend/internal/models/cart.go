package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	SessionID uuid.UUID  `db:"session_id" json:"session_id"`
	CartItems []CartItem `json:"cart_items"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

type CartItem struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CartID    uuid.UUID `db:"cart_id" json:"cart_id"`
	ProductID string    `db:"product_id" json:"product_id"`
	Product   Product   `json:"product"`
	Quantity  int       `db:"quantity" json:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
