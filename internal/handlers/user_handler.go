package handlers

import (
	"encoding/json"
	"gowhat/internal/models"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler GetUser Called")
	var userResponse models.User

	if err := json.NewDecoder(r.Body).Decode(&userResponse); err != nil {
		log.Println("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		log.Println("Encode error")
		return
	}
}
