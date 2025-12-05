package dto

import (
	"crud-service/domain/entities"
)

// request

type CreateExampleRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Count int    `json:"count" binding:"required,min=0"`
}

type UpdateExampleRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Count int    `json:"count" binding:"required,min=0"`
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
