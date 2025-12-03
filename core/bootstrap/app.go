package bootstrap

import "log/slog"

type App struct {
	config *Config
	server *Server
	logger *slog.Logger
}

func (a *App) Run() error {
	return a.server.run()
}

func NewApp(config *Config, server *Server, logger *slog.Logger) *App {
	return &App{
		config: config,
		server: server,
		logger: logger,
	}
}
