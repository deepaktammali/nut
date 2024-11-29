package handlers

import (
	"fmt"
	"net/http"
	"nut/internal/dtos"
	"nut/internal/helpers"
	"strings"
)

func buildErrorMessageFromValidationErrors(errors map[string]string) string {
	var messageBuilder strings.Builder

	for _, message := range errors {
		messageBuilder.Write([]byte(fmt.Sprintf("%s.", message)))
	}

	return messageBuilder.String()
}

func (h *Handler) createTicket(w http.ResponseWriter, r *http.Request) {
	var createTicketDto dtos.CreateTicketDto
	err := helpers.ReadJsonFromRequest(r, &createTicketDto)

	if err != nil {
		fmt.Printf("Error reading request body. Error - %s.", err)
		helpers.WriteErrorToResponse(w, "Error reading payload from request", nil, http.StatusBadRequest)
		return
	}

	validationErrors := createTicketDto.Validate()

	if len(validationErrors) > 0 {
		helpers.WriteErrorToResponse(w, buildErrorMessageFromValidationErrors(validationErrors), nil, http.StatusBadRequest)
		return
	}

	newTicket, err := h.store.Tickets.CreateTicket(createTicketDto)

	if err != nil {
		fmt.Printf(err.Error())
		helpers.WriteErrorToResponse(w, err.Error(), nil, http.StatusInternalServerError)
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

	helpers.WriteSuccessToResponse(w, ticketDto, nil, http.StatusOK)
}
