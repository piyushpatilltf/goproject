package routers

import (
	"go-crud-api/controllers"

	"github.com/gin-gonic/gin"
)

func LogRoutes(router *gin.Engine) {
	logRoutes := router.Group("/logs")
	{
		logRoutes.POST("", controllers.CreateLog)
		logRoutes.GET("", controllers.GetLogs)
		logRoutes.GET("/:id", controllers.GetLog)
		logRoutes.PUT("/:id", controllers.UpdateLog)
		logRoutes.DELETE("/:id", controllers.DeleteLog)
	}
}