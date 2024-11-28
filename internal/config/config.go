package config

import (
	"database/sql"
)

type AppConfig struct {
	Db          *sql.DB
	Environment string
}

func NewAppConfig(db *sql.DB, environment string) *AppConfig {
	config := &AppConfig{
		Db:          db,
		Environment: environment,
	}

	return config
}
