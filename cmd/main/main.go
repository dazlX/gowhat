package main

import (
	"context"
	"fmt"
	"gowhat/internal/config"
	"gowhat/internal/handlers"
	"gowhat/internal/handlers/middleware"
	"gowhat/internal/handlers/usershandler"
	"gowhat/internal/logger"
	"gowhat/internal/repository"
	"gowhat/internal/service/users"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"go.uber.org/zap"
)

func main() {

	log, logClose, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer logClose()

	log.Info("Server Start")

	cfg := config.Load()

	db, err := repository.InitialDb(cfg)
	if err != nil {
		log.Error(err.Error())
		return
	} else {
		log.Info("База данных подключена")
	}

	repo := repository.NewUserRepository(db)
	service := users.NewService(repo)
	hCfg := usershandler.NewHandlerConfig(log, service)
	// usershandler.NewHandlerConfig(log, service)
	userSwitch := usershandler.Switch(hCfg)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", middleware.Logger(handlers.Health, log))
	mux.HandleFunc("/api/users", middleware.Logger(userSwitch, log))
	port := ":8080"

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("Ошибка запуска сервера: ", zap.Error(err))
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	log.Warn("Получен сигнал на завершение сервера: ", zap.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Ошибка остановки сервера", zap.Error(err))
		return
	} else {
		log.Info("Сервер остановлен корректно")
	}
}
