package repository

import (
	"Tulupou"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateClient(client Tulupou.Client) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (username, email, phone_number, role, password) values ($1, $2, $3, $4, $5) RETURNING id", clientTable)

	row := r.db.QueryRow(q, client.Username, client.Email, client.Phone_Number, client.Role, client.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetClient(username, password string) (Tulupou.Client, error) {
	var client Tulupou.Client
	q := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 and password=$2", clientTable)
	err := r.db.Get(&client, q, username, password)

	return client, err
}
