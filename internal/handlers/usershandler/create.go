package usershandler

import (
	"encoding/json"
	"fmt"
	"gowhat/internal/models"
	"gowhat/internal/utils"
	"net/http"

	"go.uber.org/zap"
)

func (h *HandlerConfig) Create(w http.ResponseWriter, r *http.Request) {
	var usersData models.User
	log := h.Logger

	if err := json.NewDecoder(r.Body).Decode(&usersData); err != nil {
		log.Warn("Ошибка декодирования тела: ", zap.Error(err))
		utils.BadRequest(w)
		return
	}

	defer r.Body.Close()

	if err := h.Service.Create(usersData); err != nil {
		log.Error(err.Error())
		utils.BadRequest(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseMessage := fmt.Sprintf(`{"success": true, "user_data": %+v}`, usersData)

	w.Write([]byte(responseMessage))
}
