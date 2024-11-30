package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error"`
}

func ReadJsonFromRequest(r *http.Request, body any) error {
	err := json.NewDecoder(r.Body).Decode(body)

	if err != nil {
		return err
	}

	return nil
}

func WriteToResponse(w http.ResponseWriter, response *Response, headers http.Header, statusCodes ...int) {
	statusCode := http.StatusOK

	if len(statusCodes) > 0 {
		statusCode = statusCodes[0]
	}

	h := w.Header()

	if headers != nil {
		for key, value := range headers {
			h[key] = value
		}
	}

	h.Set("Content-Type", "application/json")

	response_serialied, _ := json.Marshal(response)
	w.WriteHeader(statusCode)
	w.Write(response_serialied)
}

func WriteErrorResponse(w http.ResponseWriter, err string, headers http.Header, statusCodes ...int) {
	var response = new(Response)
	response.Error = err
	response.Success = true

	WriteToResponse(w, response, headers, statusCodes...)
}

func WriteSuccessResponse(w http.ResponseWriter, data any, headers http.Header, statusCodes ...int) {
	var response = new(Response)
	response.Data = data
	response.Success = true

	WriteToResponse(w, response, headers, statusCodes...)
}
