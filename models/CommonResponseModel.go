package models

// ErrorResponse is a struct to be used for response of error
type ErrorResponse struct {
	Status  int    `json:"error_status"`
	Message string `json:"error_message"`
}

// OkResponse is a struct to be used for response of ok
type OkResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
