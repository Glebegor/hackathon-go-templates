package routers

import (
	"github.com/Glebegor/hackathon-go-templates/api/middleware"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}

func NewRouter(deps *Dependencies, env string) *gin.Engine {
	var r *gin.Engine

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	} else {
		r = gin.Default()
	}

	r.Use(middleware.CORS())

	return r
}
