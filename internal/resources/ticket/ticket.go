package ticket

type Ticket struct {
	id string
}

type TicketStore interface {
	CreateTicket(Ticket Ticket) (string, error)
	GetTicket(id string) (*Ticket, error)
	ArchiveTicket(id string) (bool, error)
	ListTickets() ([]Ticket, error)
}
