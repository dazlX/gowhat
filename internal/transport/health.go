package transport

import (
	"gowhat/internal/utils"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if !utils.GETChecked(w, r) {
		return
	}
	w.Header().Set("C`ontent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": true}`))
}
