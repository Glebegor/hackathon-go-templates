package handlers

import (
	"log/slog"
	"net/http"

	"crud-service/api/dto"
	"crud-service/bootstrap"
	"crud-service/domain/commons/interfaces"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	usecase interfaces.IExampleUsecase
	logger  *slog.Logger
	config  *bootstrap.Config
}

func NewExampleHandler(config *bootstrap.Config, usecase interfaces.IExampleUsecase, logger *slog.Logger) interfaces.IExampleHandler {
	return &ExampleHandler{
		config:  config,
		usecase: usecase,
		logger:  logger,
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

func (h *ExampleHandler) GetByID(c *gin.Context) {
	h.logger.Info("GetByID called")

	// Implementation here

	c.JSON(http.StatusOK, dto.GetExamplesResponse{
		Examples: nil,
	})
}
func (h *ExampleHandler) CreateExample(c *gin.Context) {
	h.logger.Info("CreateExample called")

	// Implementation here

	c.JSON(http.StatusOK, dto.GetExamplesResponse{
		Examples: nil,
	})
}
func (h *ExampleHandler) UpdateExample(c *gin.Context) {
	h.logger.Info("UpdateExample called")

	// Implementation here

	c.JSON(http.StatusOK, dto.GetExamplesResponse{
		Examples: nil,
	})
}
func (h *ExampleHandler) DeleteExample(c *gin.Context) {
	h.logger.Info("DeleteExample called")

	// Implementation here

	c.JSON(http.StatusOK, dto.GetExamplesResponse{
		Examples: nil,
	})
}
