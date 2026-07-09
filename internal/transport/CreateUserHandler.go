package transport

import (
	"encoding/json"
	"gowhat/internal/models"
	"net/http"

	"go.uber.org/zap"
)

func (h *HandlerConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var usersData models.User

	if err := json.NewDecoder(r.Body).Decode(&usersData); err != nil {
		h.Logger.Warn("Ошибка декодирования тела: ", zap.Error(err))

		json.NewEncoder(w).Encode(models.BadRequest{
			Success: "false",
			Error:   "Invalid JSON",
		})
		return
	}

	if err := h.Repo.CreateUser(usersData); err != nil {
		h.Logger.Warn(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		`{"success": true}`))
}
