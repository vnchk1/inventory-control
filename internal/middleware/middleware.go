package middleware

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

func LoggingMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			defer func() {
				latency := time.Since(start)
				logger.Info("completed request",
					"method", c.Request().Method,
					"path", c.Request().URL.Path,
					"status", c.Response().Status,
					"latency", latency.Milliseconds(),
					"ip", c.RealIP())
			}()
			return err
		}
	}
}
