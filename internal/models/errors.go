package models

import (
	"errors"
	"fmt"
)

var (
	ErrUnique   = errors.New("already exists")
	ErrNotFound = errors.New("not found")
	ErrEnvLoad  = errors.New("error loading .env file")
	ErrCfgPath  = errors.New("CONFIG_PATH is required")
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
