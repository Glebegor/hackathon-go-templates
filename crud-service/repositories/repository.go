package repository

import (
	"log/slog"

	"crud-service/bootstrap"
	"crud-service/domain/commons/interfaces"

	"gorm.io/gorm"
)

type Repositoryes struct {
	ExampleRepo interfaces.IExampleRepository
}

func NewRepositories(db *gorm.DB, config *bootstrap.Config, logger *slog.Logger) *Repositoryes {
	exampleRepo := NewExampleRepository(db, logger)

	return &Repositoryes{
		ExampleRepo: exampleRepo,
	}
}
