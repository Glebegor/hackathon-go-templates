package handlers

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/Glebegor/hackathon-go-templates/domain/commons/interfaces"
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
	h.logger.Info("CheckHealth called")

	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
