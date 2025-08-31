package models

import (
	"errors"
	"fmt"
)

var (
	ErrUnique        = errors.New("already exists")
	ErrExists        = errors.New("does not exists")
	ErrNotFound      = errors.New("not found")
	ErrEnvLoad       = errors.New("error loading .env file")
	ErrCfgPath       = errors.New("CONFIG_PATH is required")
	ErrFieldRequired = errors.New("field is required")
	ErrNegative      = errors.New("cannot be negative")
)

func NewErrNotFound(name string, value any) error {
	return fmt.Errorf("%s: %v, %w", name, value, ErrNotFound)
}

func NewEmptyErr(field string) error {
	return fmt.Errorf("%w: %s", ErrFieldRequired, field)
}

func NewNegativeErr(field string) error {
	return fmt.Errorf("%w: %s", ErrNegative, field)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
