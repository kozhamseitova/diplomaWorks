package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password, phone_number, role) values ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
