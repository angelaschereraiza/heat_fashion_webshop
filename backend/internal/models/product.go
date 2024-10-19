package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name"`
	Price float64   `db:"price" json:"price"`
}
