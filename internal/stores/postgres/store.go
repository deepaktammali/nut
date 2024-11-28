package postgres

import (
	"database/sql"
	"nut/internal/stores"
)

func NewStore(db *sql.DB) *stores.Store {
	ticketStore := &PostgresTicketStore{
		Db: db,
	}

	return &stores.Store{
		Tickets: ticketStore,
	}
}
