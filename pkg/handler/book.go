package handler

import (
	"Tulupou/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (h *Handler) CreateBook(c *gin.Context) {

	type AddBookRequestBody struct {
		Name string `json:"name" binding:"required"`
	}

	body := AddBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Name = body.Name

	c.JSON(http.StatusCreated, &book)
}

func (h *Handler) getAllBook(c *gin.Context) {

}

func (h *Handler) getBookById(c *gin.Context) {

}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteBook(c *gin.Context) {

}
