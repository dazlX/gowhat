package repository

import (
	"fmt"
)

func (u *UserRepository) Delete(userPhoneNumber string) error {
	result, err := u.db.Exec(
		"DELETE FROM users WHERE phone_number=$1;", userPhoneNumber)
	if err != nil {
		return fmt.Errorf("Ошибка удаления пользователя: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Ошибка получения количества удаленных строк: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("Пользователь с номером %s не найден", userPhoneNumber)
	}

	return nil
}
