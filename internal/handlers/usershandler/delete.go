package usershandler

import (
	"fmt"
	"gowhat/internal/models"
	"gowhat/internal/utils"
	"net/http"
)

func (h *HandlerConfig) Delete(w http.ResponseWriter, r *http.Request) {
	var UserData models.User
	log := h.Logger
	log.Info("ЧУВАААК")
	//TODO: Дописать

	phone := r.URL.Query().Get("phone")
	if phone == "" {
		log.Info("Номер отсутствует")
		return
	}
	log.Info(phone)
	if err := h.Service.Delete(phone); err != nil {
		log.Error(err.Error())
		utils.BadRequest(w)
		return
	}

	info := fmt.Sprintf("Пользователь %s был удален", UserData.FullName)
	log.Info(info)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	{
		"success": true,
		"message": "Пользователь был успешно удалён"
	}`))
}
