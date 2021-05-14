package models

type Order struct {
	ID    int `json:"id"    validate:"required"`
	Price int `json:"price" validate:"required"`
}
