package handlers

import (
	"log/slog"

	"crud-service/bootstrap"
	"crud-service/domain/commons/interfaces"
	usecase "crud-service/usecases"
)

type Handlers struct {
	ExampleHandler interfaces.IExampleHandler
	HealthHandler  interfaces.IHealthHandler
}

func NewHandlers(config *bootstrap.Config, logger *slog.Logger, usecases *usecase.Usecases) *Handlers {
	exampleHandler := NewExampleHandler(config, usecases.Example, logger)
	healthHandler := NewHealthHandler(config, logger)

	return &Handlers{
		ExampleHandler: exampleHandler,
		HealthHandler:  healthHandler,
	}
}
