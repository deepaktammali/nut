package dtos

func (dto *CreateTicketDto) Validate() map[string]string {
	var errors = make(map[string]string)

	if dto.Title == "" {
		errors["title"] = "Title is required field"
	}

	if dto.Description == "" {
		errors["description"] = "Description is required field"
	}

	if dto.Priority == "" {
		errors["priority"] = "Priority is required field"
	}

	return errors
}
