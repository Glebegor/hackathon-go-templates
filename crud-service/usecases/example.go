package usecase

import (
	"crud-service/domain/commons/interfaces"
	"crud-service/domain/entities"
)

type ExampleUsecase struct {
	repo interfaces.IExampleRepository
}

func NewExampleUsecase(repo interfaces.IExampleRepository) interfaces.IExampleUsecase {
	return &ExampleUsecase{
		repo: repo,
	}
}

// Methods

func (u *ExampleUsecase) GetAllExamples() ([]entities.ExampleEntity, error) {
	return u.repo.GetAllExamples()
}
func (u *ExampleUsecase) GetByID(id int) (entities.ExampleEntity, error) {
	return u.repo.GetByID(id)
}
func (u *ExampleUsecase) CreateExample(example entities.ExampleEntity) error {
	return u.repo.CreateExample(example)
}
func (u *ExampleUsecase) UpdateExample(example entities.ExampleEntity) (int, error) {
	return u.repo.UpdateExample(example)
}
func (u *ExampleUsecase) DeleteExample(id int) error {
	return u.repo.DeleteExample(id)
}
