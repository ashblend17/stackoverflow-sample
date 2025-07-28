package routes

import (
	"github.com/ashblend17/stackoverflow-sample/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "Server running."})
		})
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}
}
