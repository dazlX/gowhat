package repository

import (
	"database/sql"
	"fmt"
	"gowhat/internal/models"

	_ "github.com/lib/pq"
)

func InitialDb(cfg models.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", createConnectionString(cfg))
	if err != nil {
		return nil, fmt.Errorf("Ошибка инициализации базы данных: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("База данных не отвечает: %w", err)
	}

	return db, nil
}

func createConnectionString(cfg models.Config) string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.UserDB, cfg.PasswordDB, cfg.HostDB, cfg.PortDB, cfg.DataBaseName, cfg.SslMode)
	return connStr
}
