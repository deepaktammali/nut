package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"nut/internal/dtos"
	"nut/internal/helpers"
	"nut/internal/stores"
	"strconv"
)

func (h *Handler) createTicket(w http.ResponseWriter, r *http.Request) {
	var createTicketDto dtos.CreateTicketDto
	err := helpers.ReadJsonFromRequest(r, &createTicketDto)

	if err != nil {
		fmt.Printf("Error reading request body. Error - %s.", err)
		helpers.WriteErrorResponse(w, "Error reading payload from request", nil, http.StatusBadRequest)
		return
	}

	validationError := createTicketDto.Validate()

	if validationError != "" {
		helpers.WriteErrorResponse(w, validationError, nil, http.StatusBadRequest)
		return
	}

	newTicket, err := h.store.Tickets.CreateTicket(createTicketDto)

	if err != nil {
		fmt.Printf("Error creating ticket. Error - %s", err.Error())
		helpers.WriteErrorResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	ticketDto := &dtos.TicketDto{
		Id:          newTicket.Id,
		Title:       newTicket.Title,
		Description: newTicket.Description,
		Status:      newTicket.Status,
		Priority:    newTicket.Priority,
		CreatedAt:   newTicket.CreatedAt,
		UpdatedAt:   newTicket.UpdatedAt,
	}

	helpers.WriteSuccessResponse(w, ticketDto, nil, http.StatusOK)
}

func (h *Handler) getTicket(w http.ResponseWriter, r *http.Request) {
	ticketIdParam := r.PathValue("ticketId")
	ticketId, err := strconv.Atoi(ticketIdParam)

	if err != nil {
		helpers.WriteErrorResponse(w, fmt.Sprintf("Error parsing ticket id - passed %s", ticketIdParam), nil, http.StatusBadRequest)
		return
	}

	ticket, err := h.store.Tickets.GetTicket(ticketId)

	if err != nil {
		msg := fmt.Sprintf("Error getting ticket with id %d", ticketId)
		statusCode := http.StatusInternalServerError

		if errors.Is(err, stores.ErrTicketNotFound) {
			msg = fmt.Sprintf("Ticket with id %d not found", ticketId)
			statusCode = http.StatusNotFound
		}

		helpers.WriteErrorResponse(w, msg, nil, statusCode)
		return
	}

	ticketDto := dtos.GetTicketDtoFromTicketEntity(ticket)
	helpers.WriteSuccessResponse(w, &ticketDto, nil, http.StatusOK)
}

func (h *Handler) updateTicket(w http.ResponseWriter, r *http.Request) {
	ticketIdParam := r.PathValue("ticketId")
	ticketId, err := strconv.Atoi(ticketIdParam)

	if err != nil {
		helpers.WriteErrorResponse(w, fmt.Sprintf("Error parsing ticket id - passed %s", ticketIdParam), nil, http.StatusBadRequest)
		return
	}

	var updateTicketDto dtos.UpdateTicketDto
	err = helpers.ReadJsonFromRequest(r, &updateTicketDto)

	if err != nil {
		fmt.Printf("Error reading request body. Error - %s.", err)
		helpers.WriteErrorResponse(w, "Error reading payload from request", nil, http.StatusBadRequest)
		return
	}

	ticket, err := h.store.Tickets.UpdateTicket(ticketId, &updateTicketDto)

	if err != nil {
		msg := fmt.Sprintf("Error updating ticket with id %d", ticketId)
		statusCode := http.StatusInternalServerError

		if errors.Is(err, stores.ErrTicketNotFound) {
			msg = fmt.Sprintf("Ticket with id %d not found", ticketId)
			statusCode = http.StatusNotFound
		}

		helpers.WriteErrorResponse(w, msg, nil, statusCode)
		return
	}

	ticketDto := dtos.GetTicketDtoFromTicketEntity(ticket)
	helpers.WriteSuccessResponse(w, &ticketDto, nil, http.StatusOK)
}

func (h *Handler) listTickets(w http.ResponseWriter, _ *http.Request) {
	tickets, err := h.store.Tickets.ListTickets()

	if err != nil {
		helpers.WriteErrorResponse(w, "Error listing tickets", nil, http.StatusInternalServerError)
		return
	}

	ticketDtos := make([]*dtos.TicketDto, len(tickets))

	for _, ticket := range tickets {
		ticketDto := dtos.GetTicketDtoFromTicketEntity(ticket)
		ticketDtos = append(ticketDtos, ticketDto)
	}

	helpers.WriteSuccessResponse(w, ticketDtos, nil, http.StatusOK)
}
