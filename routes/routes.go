package routes

import (
	"github.com/ashblend17/stackoverflow-sample/controllers"
	"github.com/ashblend17/stackoverflow-sample/middlewares"

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

		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.POST("/test", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "Auth passed"}) })
			auth.POST("/questions", controllers.CreateQuestion)
			auth.POST("/questions/:id/answers", controllers.CreateAnswer)
			auth.POST("/questions/:id/vote", controllers.VoteHandler("question"))
			auth.POST("/answers/:id/vote", controllers.VoteHandler("answer"))
			api.GET("/questions/:id/summary", controllers.SummarizeQuestion)
		}
	}
}
