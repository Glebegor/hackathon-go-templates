package handlers

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/Glebegor/hackathon-go-templates/domain/commons/interfaces"
	usecase "github.com/Glebegor/hackathon-go-templates/usecases"
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
