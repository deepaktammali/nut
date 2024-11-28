package postgres

import (
	"database/sql"
	"nut/internal/resources/ticket"
)

type PostgresTicketStore struct {
	Db *sql.DB
}

func NewPostgresTaskStore(db *sql.DB) ticket.TicketStore {
	pgTaskStore := new(PostgresTicketStore)
	pgTaskStore.Db = db
	return pgTaskStore
}

func (store *PostgresTicketStore) CreateTicket(ticket ticket.Ticket) (string, error) {
	return "1", nil
}

func (store *PostgresTicketStore) GetTicket(id string) (*ticket.Ticket, error) {
	return new(ticket.Ticket), nil
}

func (store *PostgresTicketStore) ArchiveTicket(id string) (bool, error) {
	return true, nil
}

func (store *PostgresTicketStore) ListTickets() ([]ticket.Ticket, error) {
	return nil, nil
}
