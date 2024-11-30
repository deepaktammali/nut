package dtos

func (dto *CreateTicketDto) Validate() string {
	if dto.Title == "" {
		return "Title is required field"
	}

	if dto.Description == "" {
		return "Description is required field"
	}

	if dto.Priority == "" {
		return "Priority is required field"
	}

	return ""
}
