package routers

import (
	"crud-service/api/handlers"
	"crud-service/api/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "crud-service/api/docs"
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

	root := r.Group("/")
	{

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		root.GET("/health", handlers.HealthHandler.CheckHealth)

		api := root.Group("/api")
		{
			v1 := api.Group("/v2")
			{
				example := v1.Group("/examples")
				{
					example.GET("/", handlers.ExampleHandler.GetAllExamples)
					example.GET("/:id", handlers.ExampleHandler.GetByID)
					example.POST("/", handlers.ExampleHandler.CreateExample)
					example.PUT("/:id", handlers.ExampleHandler.UpdateExample)
					example.DELETE("/:id", handlers.ExampleHandler.DeleteExample)
				}
			}
		}

	}

	r.Use(middleware.CORS())

	return r
}
