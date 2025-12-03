package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Glebegor/hackathon-go-templates/api/dto"
	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/Glebegor/hackathon-go-templates/domain/interfaces"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	usecase *interfaces.ExampleUsecase
	logger  *slog.Logger
	config  *bootstrap.Config
}

func NewExampleHandler(config *bootstrap.Config, usecase *interfaces.ExampleUsecase, logger *slog.Logger) *ExampleHandler {
	return &ExampleHandler{
		config:  config,
		usecase: usecase,
		logger:  logger,
	}
}

func (h *ExampleHandler) SetupHandlers(rg *gin.RouterGroup) {
	exampleGroup := rg.Group("/examples")
	{
		exampleGroup.GET("/", h.GetAllExamples)
	}
}

// Methods
func (h *ExampleHandler) GetAllExamples(c *gin.Context) {
	h.logger.Info("GetExamples called")

	// Implementation here

	c.JSON(http.StatusOK, dto.GetExamplesResponse{
		Examples: nil,
	})
}
