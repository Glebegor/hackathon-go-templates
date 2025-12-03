package handlers

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	logger *slog.Logger
	config *bootstrap.Config
}

func NewHealthHandler(config *bootstrap.Config, logger *slog.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
		config: config,
	}
}

func (h *HealthHandler) SetupHandlers(rg *gin.RouterGroup) {
	healthGroup := rg.Group("/health")
	{
		healthGroup.GET("/", h.CheckHealth)
	}
}

// Methods
func (h *HealthHandler) CheckHealth(c *gin.Context) {
	h.logger.Info("CheckHealth called")

	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
