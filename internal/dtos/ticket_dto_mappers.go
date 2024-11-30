package dtos

import "nut/internal/entities"

func GetTicketDtoFromTicketEntity(ticket *entities.Ticket) *TicketDto {
	dto := &TicketDto{
		Id:          ticket.Id,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      ticket.Status,
		Priority:    ticket.Priority,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}

	return dto
}
