package server

import (
	"log/slog"
)

type Handlers struct {
	Products *ProductHandler
}

func NewHandlers(products ProductService, logger *slog.Logger) *Handlers {
	return &Handlers{
		Products: NewProductHandler(products, logger),
	}
}
