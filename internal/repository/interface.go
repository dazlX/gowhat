package repository

import (
	"database/sql"
	"gowhat/internal/models"
)

type RepositoryInterface interface {
	CreateUser(db *sql.DB, user models.User) (bool, error)
	UpdateUser()
	DeleteUser()
}
