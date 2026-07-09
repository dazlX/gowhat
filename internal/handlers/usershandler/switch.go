package usershandler

import (
	"net/http"
)

func Switch(hCfg *HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			//Создать пользователя
			hCfg.Create(w, r)
		case http.MethodGet:
			//Запросить пользователя
		case http.MethodDelete:
			//TODO: Удалить пользователя
			hCfg.Delete(w, r)
		case http.MethodPatch:
		}
	}
}
