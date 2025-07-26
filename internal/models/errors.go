package models

const (
	InvalidRequestBodyString  = "Invalid request body"
	InternalServerErrorString = "Internal server error"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
