package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"nut/internal/constants"
	"nut/internal/dtos"
	"nut/internal/entities"
	"nut/internal/helpers"
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
	var ticket = entities.Ticket{}
	err := store.Db.QueryRow("SELECT id, title, description, status, priority, created_at, updated_at FROM tickets WHERE id=$1", id).Scan(&ticket.Id, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.Priority, &ticket.CreatedAt, &ticket.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, stores.ErrTicketNotFound
		}

		return nil, errors.New(fmt.Sprintf("Error getting ticket with id %d", id))
	}

	return &ticket, nil
}

func (store *PostgresTicketStore) UpdateTicket(id int, updateTicketDto *dtos.UpdateTicketDto) (*entities.Ticket, error) {
	dbTicket, err := store.GetTicket(id)

	if err != nil {
		return nil, err
	}

	dbTicket.Title = helpers.FirstNonEmpty(updateTicketDto.Title, dbTicket.Title)
	dbTicket.Description = helpers.FirstNonEmpty(updateTicketDto.Description, dbTicket.Description)
	dbTicket.Status = helpers.FirstNonEmpty(updateTicketDto.Status, dbTicket.Status)
	dbTicket.Priority = helpers.FirstNonEmpty(updateTicketDto.Priority, dbTicket.Priority)
	dbTicket.UpdatedAt = time.Now().UTC()

	query := `UPDATE tickets
	SET title=$2, description=$3, status=$4, priority=$5, updated_at=$6
	WHERE id = $1
	`
	_, err =
		store.Db.Exec(query, id, dbTicket.Title, dbTicket.Description, dbTicket.Status, dbTicket.Priority, dbTicket.UpdatedAt.Format(time.DateTime))

	if err != nil {
		log.Printf("Error updating ticket with id %d. Err - %s.", id, err.Error())
		return nil, errors.New(fmt.Sprintf("Error updating ticket with id %d", id))
	}

	return dbTicket, nil
}

func (store *PostgresTicketStore) ListTickets() ([]*entities.Ticket, error) {
	query := "SELECT id, title, description, status, priority, created_at, updated_at FROM tickets ORDER BY id"
	rows, err := store.Db.Query(query)

	if err != nil {
		return nil, errors.New("Error getting list of all tickets")
	}

	defer rows.Close()

	var tickets []*entities.Ticket

	for rows.Next() {
		ticket := &entities.Ticket{}
		err := rows.Scan(&ticket.Id, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.Priority, &ticket.CreatedAt, &ticket.UpdatedAt)

		if err != nil {
			return tickets, err
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return tickets, err
	}

	return tickets, nil
}
