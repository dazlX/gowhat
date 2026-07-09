package users

import (
	"gowhat/internal/models"
	"gowhat/internal/repository"
)

type ServiceInterface interface {
	Create(user models.User) error
	Get()
	Delete(user models.User) error
	Update()
}

type Service struct {
	repo *repository.UserRepository
}

func NewService(repo *repository.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}
