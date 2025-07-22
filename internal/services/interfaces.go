package repo

import (
	"github.com/jackc/pgx/v5"
	"github.com/vnchk1/inventory-control/internal/models"
	"log/slog"
)

type (
	ProductCrud interface {
		Create(*pgx.Conn, *slog.Logger, *models.Product) error
		Read(*pgx.Conn, *slog.Logger, int) error
		Update(*pgx.Conn, *slog.Logger, *models.Product) error
		Delete(*pgx.Conn, *slog.Logger, int) error
	}
)
