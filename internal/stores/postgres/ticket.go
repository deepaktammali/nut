package postgres

import (
	"database/sql"
	"fmt"
	"nut/internal/constants"
	"nut/internal/dtos"
	"nut/internal/entities"
	"nut/internal/stores"
	"time"
)

type PostgresTicketStore struct {
	Db *sql.DB
}

func NewPostgresTaskStore(db *sql.DB) stores.TicketStore {
	pgTaskStore := new(PostgresTicketStore)
	pgTaskStore.Db = db
	return pgTaskStore
}

func (store *PostgresTicketStore) CreateTicket(createTicketDto dtos.CreateTicketDto) (*entities.Ticket, error) {
	var ticketId int

	utc_now := time.Now().UTC()
	formatted_utc_now := utc_now.Format(time.DateTime)
	err := store.Db.QueryRow("INSERT INTO tickets (title, description, status, priority, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;",
		createTicketDto.Title, createTicketDto.Description, constants.TicketStatusOpen, createTicketDto.Priority, formatted_utc_now, formatted_utc_now).Scan(&ticketId)

	if err != nil {
		return new(entities.Ticket), fmt.Errorf("Cannot create ticket. Error - %s", err)
	}

	newTicket := entities.Ticket{
		Id:          ticketId,
		Title:       createTicketDto.Title,
		Description: createTicketDto.Description,
		Status:      constants.TicketStatusOpen,
		Priority:    createTicketDto.Priority,
		CreatedAt:   utc_now,
		UpdatedAt:   utc_now,
	}

	return &newTicket, nil
}

func (store *PostgresTicketStore) GetTicket(id int) (*entities.Ticket, error) {
	return new(entities.Ticket), nil
}

func (store *PostgresTicketStore) ArchiveTicket(id int) (bool, error) {
	return true, nil
}

func (store *PostgresTicketStore) ListTickets() ([]entities.Ticket, error) {
	return nil, nil
}
