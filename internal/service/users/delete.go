package users

import "fmt"

func (s *Service) Delete(phone string) error {
	if phone == "" {
		return fmt.Errorf("Неправильный формат номера")
	}
	return s.repo.Delete(phone)

}
