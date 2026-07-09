package models

type User struct {
	ID int `json:"id"`
	FullName string `json:"full_name"`
	Age int `json:"age"`
	PhoneNumber *string `json:"phone_number"`
}