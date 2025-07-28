package models

import (
	"errors"
	"fmt"
)

var (
	ErrUnique                = errors.New("already exists")
	ErrNotFound              = errors.New("not found")
	ErrNotUpdated            = errors.New("not updated")
	ErrConfigPathNotProvided = errors.New("config path didn't provide")
	ErrBadConfigPort         = errors.New("port must be upper than 0")
	ErrBadResponseTime       = errors.New("response time must be upper than 0ms")
	ErrBadUserName           = errors.New("username is required")
	ErrBadPassword           = errors.New("password is required")
	ErrBadDBName             = errors.New("db name is required")
	ErrBadLogLevel           = errors.New("log level is required")
	ErrMigrationsNotProvided = errors.New("migrations path didn't provide")
	ErrBadHost               = errors.New("host is required")
	ErrBadPort               = errors.New("port is required")
	ErrBadSSLMode            = errors.New("SSL_MODE is required")
	ErrBadRequestBody        = errors.New("bad request body")
)

func NewEmptyErr(field string) error {
	return fmt.Errorf("field '%s' is required", field)
}

func NewNegativeErr(field string) error {
	return fmt.Errorf("field '%s' cannot be negative", field)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
