package routes

import (
	"github.com/fullstacktf/fitness-backend/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter function
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controller.GetWelcome)
	}
	return r
}
