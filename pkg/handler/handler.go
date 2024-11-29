package handler

import (
	"Tulupou/pkg/service"
	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "Tulupou/docs"
)

type Handler struct {
	services *service.Service
	db       *sqlx.DB
}

func NewHandler(services *service.Service, db *sqlx.DB) *Handler {
	return &Handler{
		services: services,
		db:       db,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.clientIdentity)
	{
		client := api.Group("/client")
		{
			client.GET("/", h.getAllClient)
			client.GET("/:id", h.getClientById)
			client.PUT("/:id", h.updateClient)
			client.DELETE("/:id", h.deleteClient)
		}

		book := api.Group("/book")
		{
			book.POST("/", h.CreateBook)
			book.GET("/", h.getAllBook)
			book.GET("/:id", h.getBookById)
			book.PUT("/:id", h.updateBook)
			book.DELETE("/:id", h.deleteBook)
		}
	}

	return router
}
