package handlers

import "nut/internal/entities"

func (h *Handler) getTicket() *entities.Ticket {
	t, _ := h.store.Tickets.GetTicket("1")
	return t
}
