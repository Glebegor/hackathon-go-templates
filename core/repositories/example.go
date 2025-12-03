package repository

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/domain/commons/interfaces"
	"github.com/Glebegor/hackathon-go-templates/domain/entities"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewExampleRepository(db *gorm.DB, logger *slog.Logger) interfaces.IExampleRepository {
	return &ExampleRepository{
		db:     db,
		logger: logger,
	}
}

// Methods
func (r *ExampleRepository) GetAllExamples() ([]*entities.ExampleEntity, error) {
	var examples []*entities.ExampleEntity
	if err := r.db.Find(&examples).Error; err != nil {
		r.logger.Error("failed to get all examples", "error", err)
		return nil, err
	}
	return examples, nil
}

func (r *ExampleRepository) GetByID(id int) (*entities.ExampleEntity, error) {
	var example entities.ExampleEntity
	if err := r.db.First(&example, id).Error; err != nil {
		r.logger.Error("failed to get example by id", "id", id, "error", err)
		return nil, err
	}
	return &example, nil
}

func (r *ExampleRepository) CreateExample(example *entities.ExampleEntity) error {
	if err := r.db.Create(example).Error; err != nil {
		r.logger.Error("failed to create example", "error", err)
		return err
	}
	return nil
}

func (r *ExampleRepository) UpdateExample(example *entities.ExampleEntity) error {
	if err := r.db.Save(example).Error; err != nil {
		r.logger.Error("failed to update example", "id", example.Id, "error", err)
		return err
	}
	return nil
}

func (r *ExampleRepository) DeleteExample(id int) error {
	if err := r.db.Delete(&entities.ExampleEntity{}, id).Error; err != nil {
		r.logger.Error("failed to delete example", "id", id, "error", err)
		return err
	}
	return nil
}
