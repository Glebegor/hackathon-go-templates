package repository

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/Glebegor/hackathon-go-templates/domain/commons/interfaces"
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
