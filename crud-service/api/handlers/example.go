package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"crud-service/api/dto"
	"crud-service/bootstrap"
	"crud-service/domain/commons/interfaces"

	"crud-service/domain/entities"

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

// GetAllExamples godoc
// @Summary      Get all examples
// @Description  Retrieves a list of example entities. Accepts an optional JSON payload to filter, sort or paginate the results.
// @Tags         Examples
// @Accept       json
// @Produce      json
// @Success      200      {object}  dto.GetAllExamplesResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /examples [get]
func (h *ExampleHandler) GetAllExamples(c *gin.Context) {

	var data []entities.ExampleEntity

	data, err := h.usecase.GetAllExamples()
	if err != nil {
		h.logger.Error("Failed to get all examples", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to get examples",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.GetAllExamplesResponse{
		Examples: data,
	})
}

// GetByID godoc
// @Summary      Get example by ID
// @Description  Retrieves a single example entity by its ID.
// @Tags         Examples
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Example ID"
// @Success      200  {object}  dto.GetByIdExampleResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /examples/{id} [get]
func (h *ExampleHandler) GetByID(c *gin.Context) {
	var data entities.ExampleEntity

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid id parameter", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid id parameter",
			Status:  http.StatusBadRequest,
		})
		return
	}

	data, err = h.usecase.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get example by ID", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to get example",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.GetByIdExampleResponse{
		Example: data,
	})
}

// CreateExample godoc
// @Summary      Create a new example
// @Description  Creates a new example entity.
// @Tags         Examples
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateExampleRequest  true  "Create example request"
// @Success      200      {object}  dto.CreateExampleResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /examples [post]
func (h *ExampleHandler) CreateExample(c *gin.Context) {
	var data entities.ExampleEntity
	var r dto.CreateExampleRequest

	err := c.ShouldBindJSON(&r)
	if err != nil {
		h.logger.Error("Failed to bind CreateExampleRequest", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// map fields from DTO to entity
	data.Name = r.Name
	data.Count = r.Count

	err = h.usecase.CreateExample(data)
	if err != nil {
		h.logger.Error("Failed to create example", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to create example",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.CreateExampleResponse{
		Example: data,
	})
}

// UpdateExample godoc
// @Summary      Update an existing example
// @Description  Updates an example entity by ID.
// @Tags         Examples
// @Accept       json
// @Produce      json
// @Param        id       path      int                      true  "Example ID"
// @Param        request  body      dto.UpdateExampleRequest true  "Update example request"
// @Success      200      {object}  dto.UpdateExampleResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /examples/{id} [put]
func (h *ExampleHandler) UpdateExample(c *gin.Context) {
	var data entities.ExampleEntity
	var r dto.UpdateExampleRequest

	// parse id from path
	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid id parameter", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid id parameter",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// bind request body
	err = c.ShouldBindJSON(&r)
	if err != nil {
		h.logger.Error("Failed to bind UpdateExampleRequest", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// map DTO to entity (populate ID from path)
	data.Name = r.Name
	data.Count = r.Count

	// call usecase to update
	updatedID, err := h.usecase.UpdateExample(data)
	if err != nil {
		h.logger.Error("Failed to update example", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to update example",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.UpdateExampleResponse{
		Id: updatedID,
	})
}

// DeleteExample godoc
// @Summary      Delete an example
// @Description  Deletes an example entity by ID.
// @Tags         Examples
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Example ID"
// @Success      200  {object}  dto.DeleteExampleResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /examples/{id} [delete]
func (h *ExampleHandler) DeleteExample(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid id parameter", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid id parameter",
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = h.usecase.DeleteExample(id)
	if err != nil {
		h.logger.Error("Failed to delete example", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to delete example",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.DeleteExampleResponse{
		Id: id,
	})
}
