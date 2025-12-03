package routers

import (
	"github.com/Glebegor/hackathon-go-templates/api/handlers"
	"github.com/Glebegor/hackathon-go-templates/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(handlers *handlers.Handlers, env string) *gin.Engine {
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
	{
		root.GET("/health", handlers.HealthHandler.CheckHealth)
		api := root.Group("/api")
		{
			example := api.Group("/examples")
			{
				example.GET("/", handlers.ExampleHandler.GetAllExamples)
				example.GET("/:id", handlers.ExampleHandler.GetByID)
				example.POST("/", handlers.ExampleHandler.CreateExample)
				example.PUT("/:id", handlers.ExampleHandler.UpdateExample)
				example.DELETE("/:id", handlers.ExampleHandler.DeleteExample)
			}
		}

	}
	return r
}
