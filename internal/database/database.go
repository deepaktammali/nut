package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const DatabaseDriverPgx = "pgx"

func BuildPostgresDSN() string {
	envItems := []string{
		"DB_USER",
		"DB_PASS",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
	}
	dbParams := make(map[string]string)

	for _, envItem := range envItems {
		envValue, exists := os.LookupEnv(envItem)

		if !exists {
			panic(fmt.Sprintf("Environment variable %s is required to be set but not set", envItem))
		}

		if envItem == "" {
			panic(fmt.Sprintf("Environment variable %s is required to be non empty", envItem))
		}

		dbParams[envItem] = envValue
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		dbParams["DB_USER"], dbParams["DB_PASS"], dbParams["DB_HOST"], dbParams["DB_PORT"], dbParams["DB_NAME"])
}

func NewDb(driverName string, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping the database: %v\n", err)
		return nil, err
	}

	return db, nil
}
