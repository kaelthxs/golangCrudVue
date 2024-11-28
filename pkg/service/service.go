package service

import (
	"Tulupou"
	"Tulupou/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateClient(client Tulupou.Client) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Client
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
