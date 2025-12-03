package dto

import "crud-service/domain/entities"

// request
type GetExamplesRequest struct {
}

// response
type GetExamplesResponse struct {
	Examples []entities.ExampleEntity `json:"examples"`
}
