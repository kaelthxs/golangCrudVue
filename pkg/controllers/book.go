package controllers

import "github.com/jmoiron/sqlx"

type BookHandler struct {
	db *sqlx.DB
}
