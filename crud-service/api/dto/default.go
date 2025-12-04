package dto

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
