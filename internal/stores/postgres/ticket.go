package postgres

import (
	"database/sql"
	"nut/internal/entities"
	"nut/internal/stores"
)

type PostgresTicketStore struct {
	Db *sql.DB
}

func NewPostgresTaskStore(db *sql.DB) stores.TicketStore {
	pgTaskStore := new(PostgresTicketStore)
	pgTaskStore.Db = db
	return pgTaskStore
}

func (store *PostgresTicketStore) CreateTicket(ticket entities.Ticket) (string, error) {
	return "1", nil
}

func (store *PostgresTicketStore) GetTicket(id string) (*entities.Ticket, error) {
	return new(entities.Ticket), nil
}

func (store *PostgresTicketStore) ArchiveTicket(id string) (bool, error) {
	return true, nil
}

func (store *PostgresTicketStore) ListTickets() ([]entities.Ticket, error) {
	return nil, nil
}
