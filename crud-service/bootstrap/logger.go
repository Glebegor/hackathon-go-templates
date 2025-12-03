package bootstrap

import (
	"log/slog"
	"os"
	"strings"
)

type LoggerConfig struct {
	Env   string // dev, prod
	Level string // debug, info, warn, error
}

func NewLogger(cfg LoggerConfig) *slog.Logger {
	level := ParseLogLevel(cfg.Level)

	var handler slog.Handler
	if strings.ToLower(cfg.Env) == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	} else if strings.ToLower(cfg.Env) == "dev" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}
	return slog.New(handler)
}

func ParseLogLevel(v string) slog.Leveler {
	switch strings.ToLower(v) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
