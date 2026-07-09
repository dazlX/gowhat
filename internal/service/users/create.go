package users

import (
	"fmt"
	"gowhat/internal/models"
	"gowhat/internal/utils"
)

func (u *Service) Create(user models.User) error {
	if !utils.CheckNilStruct(user) {
		return fmt.Errorf("Invalid JSON")
	}

	return u.repo.Create(user)
}
