package dto

import "github.com/Glebegor/hackathon-go-templates/domain/entities"

// request
type GetExamplesRequest struct {
}

// response
type GetExamplesResponse struct {
	Examples []entities.ExampleEntity `json:"examples"`
}
