package repository

import (
	"Tulupou"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateClient(client Tulupou.Client) (int, error)
	GetClient(username, password string) (Tulupou.Client, error)
}

type Client interface {
	GetAllClient(c *gin.Context)
	GetClientById(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeleteClient(c *gin.Context)
}

type Book interface {
	CreateBook(c *gin.Context)
}

type Repository struct {
	Authorization
	Client
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
