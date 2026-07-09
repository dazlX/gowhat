package repository

import (
	"gowhat/internal/models"
)

type RepositoryInterface interface {
	Create(user models.User) error
	Delete(user models.User) error
	Update()
}
