package utils

import "net/http"

func GETChecked(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Only GET allowed!"}`))
		return false
	}
	return true
}
