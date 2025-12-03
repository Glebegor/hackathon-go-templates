package interfaces

import (
	"crud-service/domain/entities"

	"github.com/gin-gonic/gin"
)

type IExampleHandler interface {
	GetAllExamples(c *gin.Context)
	GetByID(c *gin.Context)
	CreateExample(c *gin.Context)
	UpdateExample(c *gin.Context)
	DeleteExample(c *gin.Context)
}

type IExampleRepository interface {
	GetAllExamples() ([]*entities.ExampleEntity, error)
	GetByID(id int) (*entities.ExampleEntity, error)
	CreateExample(example *entities.ExampleEntity) error
	UpdateExample(example *entities.ExampleEntity) error
	DeleteExample(id int) error
}

type IExampleUsecase interface {
	GetAllExamples() ([]*entities.ExampleEntity, error)
	GetByID(id int) (*entities.ExampleEntity, error)
	CreateExample(example *entities.ExampleEntity) error
	UpdateExample(example *entities.ExampleEntity) error
	DeleteExample(id int) error
}
