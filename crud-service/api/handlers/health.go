package handlers

import (
	"log/slog"

	"crud-service/bootstrap"
	"crud-service/domain/commons/interfaces"

	"crud-service/api/dto"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	logger *slog.Logger
	config *bootstrap.Config
}

func NewHealthHandler(config *bootstrap.Config, logger *slog.Logger) interfaces.IHealthHandler {
	return &HealthHandler{
		logger: logger,
		config: config,
	}
}

// Methods

func (h *HealthHandler) CheckHealth(c *gin.Context) {

	c.JSON(200, dto.SuccessResponse{
		Message: "Service is healthy",
		Status:  200,
	})
}
