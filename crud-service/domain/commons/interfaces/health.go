package interfaces

import "github.com/gin-gonic/gin"

type IHealthHandler interface {
	CheckHealth(c *gin.Context)
}
