package usecase

import (
	"log/slog"

	"crud-service/bootstrap"
	"crud-service/domain/commons/interfaces"
	repository "crud-service/repositories"
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
