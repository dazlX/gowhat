package utils

import "net/http"

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"success": false, "error":"Invalid JSON"}`))
}

// BUG: Починить, всегда озвращает одно и то же, разрешить писать сообщение
