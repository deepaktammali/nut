package stores

import "nut/internal/entities"

type TicketStore interface {
	CreateTicket(Ticket entities.Ticket) (string, error)
	GetTicket(id string) (*entities.Ticket, error)
	ArchiveTicket(id string) (bool, error)
	ListTickets() ([]entities.Ticket, error)
}
