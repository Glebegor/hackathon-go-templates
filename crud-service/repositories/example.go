package repository

import (
	"log/slog"

	"crud-service/domain/commons/interfaces"
	"crud-service/domain/entities"

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
func (r *ExampleRepository) GetAllExamples() ([]entities.ExampleEntity, error) {
	var examples []entities.ExampleEntity
	if err := r.db.Find(&examples).Error; err != nil {
		r.logger.Error("failed to get all examples", "error", err)
		return nil, err
	}
	return examples, nil
}

func (r *ExampleRepository) GetByID(id int) (entities.ExampleEntity, error) {
	var example entities.ExampleEntity
	if err := r.db.First(&example, id).Error; err != nil {
		r.logger.Error("failed to get example by id", "id", id, "error", err)
		return entities.ExampleEntity{}, err
	}
	return example, nil
}
func (r *ExampleRepository) CreateExample(example entities.ExampleEntity) error {
	// ensure zero value for ID so the DB can auto-increment it
	example.Id = 0

	if err := r.db.Create(&example).Error; err != nil {
		r.logger.Error("failed to create example", "error", err)
		return err
	}

	// GORM will automatically populate CreatedAt/UpdatedAt if those fields exist in the entity.
	r.logger.Info("created example", "id", example.Id)
	return nil
}

func (r *ExampleRepository) UpdateExample(example entities.ExampleEntity) (int, error) {
	if err := r.db.Save(&example).Error; err != nil {
		r.logger.Error("failed to update example", "id", example.Id, "error", err)
		return 0, err
	}
	return (int)(example.Id), nil
}

func (r *ExampleRepository) DeleteExample(id int) error {
	if err := r.db.Delete(&entities.ExampleEntity{}, id).Error; err != nil {
		r.logger.Error("failed to delete example", "id", id, "error", err)
		return err
	}
	return nil
}
