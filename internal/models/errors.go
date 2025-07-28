package models

const (
	InvalidRequestBodyString  = "Invalid request body"
	InternalServerErrorString = "Internal server error"
	InvalidIDString           = "Invalid ID"
	ProductNotFoundString     = "Product not found"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
