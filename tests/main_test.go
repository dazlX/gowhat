package tests

import (
	"gowhat/internal/config"
	"gowhat/internal/handlers"
	"gowhat/internal/logger"
	"log"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Server Start")

	log.Println("Start load config")

	cfg := config.Load()

	log.Println("Config load:", cfg)
	loge, logClose, err := logger.NewLogger()
	if err != nil {

		return
	}
	defer logClose()

	loge.Debug("Test")

	http.HandleFunc("/", handlers.GetUser)

	port := ":8080"

	if err := http.ListenAndServe(port, nil); err != nil {
		return
	}
}
