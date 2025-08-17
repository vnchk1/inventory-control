package server

import (
	"log/slog"
)

type Handlers struct {
	Categories *CategoryHandler
	Products   *ProductHandler
}

func NewHandlers(categories CategoryUseCase, products ProductUseCase, logger *slog.Logger) *Handlers {
	return &Handlers{
		Categories: NewCategoryHandler(categories, logger),
		Products:   NewProductHandler(products, logger),
	}
}
