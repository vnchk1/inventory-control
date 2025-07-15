package logger

import (
	"log/slog"
	"os"
)

func NewLogger(lvlStr string) (logger *slog.Logger) {
	lvl := ConvertLogLvl(lvlStr)
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	logger = slog.New(logHandler)
	return
}

func ConvertLogLvl(lvl string) slog.Level {
	switch lvl {
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
