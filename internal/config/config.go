package config

import (
	"database/sql"
	"nut/internal/stores"
)

type AppConfig struct {
	Db          *sql.DB
	Environment string
	store       *stores.Store
}

func NewAppConfig(db *sql.DB, environment string) *AppConfig {
	config := &AppConfig{
		Db:          db,
		Environment: environment,
	}

	return config
}
