package routers

import (
	"log/slog"

	"github.com/Glebegor/hackathon-go-templates/api/handlers"
	"github.com/Glebegor/hackathon-go-templates/api/middleware"
	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	"github.com/gin-gonic/gin"
)

type RouterDependencies struct {
	HealthHandler  *handlers.HealthHandler
	ExampleHandler *handlers.ExampleHandler
}

func NewRouterDependencies(config *bootstrap.Config, logger *slog.Logger) *RouterDependencies {
	healthHandler := handlers.NewHealthHandler(config, logger)
	exampleHandler := handlers.NewExampleHandler(config, logger)

	return &RouterDependencies{
		HealthHandler:  healthHandler,
		ExampleHandler: exampleHandler,
	}
}

func NewRouter(deps *RouterDependencies, env string) *gin.Engine {
	var r *gin.Engine

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	} else {
		r = gin.Default()
	}

	r.Use(middleware.CORS())

	root := r.Group("/")
	deps.HealthHandler.SetupHandlers(root) // Setup health routes

	api := root.Group("/api")

	deps.ExampleHandler.SetupHandlers(api) // Setup example routes

	return r
}
