package models

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrEnvLoad       = errors.New("error loading .env file")
	ErrCfgPath       = errors.New("CONFIG_PATH is required")
	ErrFieldRequired = errors.New("field is required")
	ErrNegative      = errors.New("cannot be negative")
	ErrTooManyItems  = errors.New("too many items")
)

func NewEmptyErr(field string) error {
	return fmt.Errorf("%w: %s", ErrFieldRequired, field)
}

func NewNegativeErr(field string) error {
	return fmt.Errorf("%w: %s", ErrNegative, field)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
