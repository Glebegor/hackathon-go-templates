package interfaces

import "github.com/Glebegor/hackathon-go-templates/domain/entities"

type ExampleRepository interface {
	GetAllExamples() ([]*entities.ExampleEntity, error)
	GetByID(id int) (*entities.ExampleEntity, error)
	CreateExample(example *entities.ExampleEntity) error
	UpdateExample(example *entities.ExampleEntity) error
	DeleteExample(id int) error
}

type ExampleUsecase interface {
	GetAllExamples() ([]*entities.ExampleEntity, error)
	GetByID(id int) (*entities.ExampleEntity, error)
	CreateExample(example *entities.ExampleEntity) error
	UpdateExample(example *entities.ExampleEntity) error
	DeleteExample(id int) error
}
