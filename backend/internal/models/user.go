package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Mail        string    `db:"mail" json:"mail"`
	Password    string    `db:"password" json:"-"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Address     string    `db:"address" json:"address"`
	City        string    `db:"city" json:"city"`
	PostalCode  string    `db:"postal_code" json:"postal_code"`
	Country     string    `db:"country" json:"country"`
}
