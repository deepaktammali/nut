package stores

import (
	"nut/internal/dtos"
	"nut/internal/entities"
)

type TicketStore interface {
	CreateTicket(Ticket dtos.CreateTicketDto) (*entities.Ticket, error)
	GetTicket(id int) (*entities.Ticket, error)
	ArchiveTicket(id int) (bool, error)
	ListTickets() ([]entities.Ticket, error)
}
