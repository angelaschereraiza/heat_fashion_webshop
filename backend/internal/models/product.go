package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Price float64   `db:"price"`
}
