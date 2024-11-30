package entities

import "time"

type Ticket struct {
	Id          int
	Title       string
	Description string
	Status      string
	Priority    string
	Archived    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
