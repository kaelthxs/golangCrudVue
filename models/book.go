package models

type Book struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" binding:"required"`
}
