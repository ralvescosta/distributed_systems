package models

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func StringToErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}
