package handler

import (
	"Tulupou"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

type ClientService struct {
	db *sqlx.DB
}

// @Summary Showall Tulupou.Client
// @Security ApiKeyAuth
// @Tags clients
// @ID show-clients
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /api/client/ [get]
func (r *Handler) getAllClient(c *gin.Context) {
	log.Println("getAllClient called")

	var clients []Tulupou.Client

	query := "SELECT id, username, email, phone_number, role, password FROM client"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&clients, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d clients\n", len(clients))

	if len(clients) == 0 {
		c.JSON(200, gin.H{"message": "No clients found"})
		return
	}

	c.JSON(200, clients)
}

// @Summary Showone Tulupou.Client
// @Security ApiKeyAuth
// @Tags client
// @ID show-client
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /api/client/:id [get]
func (h *Handler) getClientById(c *gin.Context) {
	// Получаем ID из параметров запроса
	id := c.Param("id")

	// Проверяем, что ID — это число
	var client Tulupou.Client
	query := "SELECT id, username, email, phone_number, role, password FROM client WHERE id = $1"

	err := h.db.Get(&client, query, id)
	if err != nil {
		log.Println("Error fetching client by ID:", err)
		c.JSON(404, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(200, client)
}

// @Summary Updateone Tulupou.Client
// @Security ApiKeyAuth
// @Tags Updateclient
// @ID Update-clients
// @Accept json
// @Produce json
// @Param input body Tulupou.Client true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /api/client/:id [put]
func (h *Handler) updateClient(c *gin.Context) {
	id := c.Param("id") // Получаем ID клиента из параметров

	var input Tulupou.Client
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE client
		SET username = $1, email = $2, phone_number = $3, role = $4, password = $5
		WHERE id = $6
	`

	_, err := h.db.Exec(query, input.Username, input.Email, input.Phone_Number, input.Role, input.Password, id)
	if err != nil {
		log.Println("Error updating client:", err)
		c.JSON(500, gin.H{"error": "Failed to update client"})
		return
	}

	c.JSON(200, gin.H{"message": "Client updated successfully"})
}

// @Summary Deleteone Tulupou.Client
// @Security ApiKeyAuth
// @Tags Deleteclient
// @ID Delete-clients
// @Accept json
// @Produce json
// @Param input body integer true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /api/client/:id [delete]
func (h *Handler) deleteClient(c *gin.Context) {
	id := c.Param("id") // Получаем ID клиента из параметров

	query := "DELETE FROM client WHERE id = $1"

	_, err := h.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting client:", err)
		c.JSON(500, gin.H{"error": "Failed to delete client"})
		return
	}

	c.JSON(200, gin.H{"message": "Client deleted successfully"})
}
