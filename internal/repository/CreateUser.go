package repository

import (
	"fmt"
	"gowhat/internal/models"
)

func (u *UserRepository) CreateUser(user models.User) error {
	_, err := u.db.Exec(
		"INSERT INTO users (full_name, age, phone_number) VALUES ($1, $2, $3);",
		user.FullName, user.Age, user.PhoneNumber)

	if err != nil {
		return fmt.Errorf("Ошибка выполнения запроса: ", err)
	}

	return nil
}
