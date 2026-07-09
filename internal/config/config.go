package config

import (
	"gowhat/internal/models"
	"gowhat/internal/utils"
)

func Load() models.Config {
	return models.Config{
		UserDB:       utils.GetEnv("USER_DB", "postgres"),
		PasswordDB:   utils.GetEnv("PASSWORD_DB", "admin"),
		PortDB:       utils.GetEnv("PORT_DB", "5432"),
		HostDB:       utils.GetEnv("HOST_DB", "localhost"),
		DataBaseName: utils.GetEnv("DATABASE_NAME", "gowhat"),
		SslMode:      utils.GetEnv("SSL_MODE", "disable"),
	}
}
