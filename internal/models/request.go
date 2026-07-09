package models

type BadRequest struct {
	Success string `json:"success"`
	Error   string `json:"error"`
}
