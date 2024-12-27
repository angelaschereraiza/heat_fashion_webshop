package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID                uuid.UUID         `db:"id" json:"id"`
	CategoryID        uuid.UUID         `db:"category_id" json:"category_id"`
	AliExpressID      int               `db:"ali_express_id" json:"ali_express_id"`
	AliExpressDetails AliExpressDetails `json:"ali_express_details"`
	Name              string            `db:"name" json:"name"`
	Price             float64           `db:"price" json:"price"`
	CreatedAt         time.Time         `db:"created_at"`
	UpdatedAt         time.Time         `db:"updated_at"`
}
