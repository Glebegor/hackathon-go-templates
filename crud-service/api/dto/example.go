package dto

import (
	"crud-service/domain/entities"
	"time"
)

// request
type CreateExampleRequest struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type UpdateExampleRequest struct {
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type DeleteExampleRequest struct {
}

// response
type GetAllExamplesResponse struct {
	Examples []entities.ExampleEntity `json:"examples"`
}

type GetByIdExampleResponse struct {
	Example entities.ExampleEntity `json:"example"`
}
type UpdateExampleResponse struct {
	Id int `json:"id"`
}
type CreateExampleResponse struct {
	Example entities.ExampleEntity `json:"example"`
}
type DeleteExampleResponse struct {
	Id int `json:"id"`
}
