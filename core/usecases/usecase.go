package usecase

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/Glebegor/hackathon-go-templates/domain/commons/interfaces"
	repository "github.com/Glebegor/hackathon-go-templates/repositories"
)

type Usecases struct {
	Example interfaces.IExampleUsecase
}

func NewUsecases(config *bootstrap.Config, logger *slog.Logger, repos *repository.Repositoryes) *Usecases {
	exampleUsecase := NewExampleUsecase(repos.ExampleRepo)

	return &Usecases{
		Example: exampleUsecase,
	}
}
