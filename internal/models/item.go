package models

type Item struct {
	ID    int    `json:"id"    validate:"required"`
	Title string `json:"title" validate:"required"`
	Price int    `json:"price" validate:"required"`
}
