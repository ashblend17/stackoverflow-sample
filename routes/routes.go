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
			auth.GET("/test", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "Auth passed"}) })
			auth.POST("/createQuestion", controllers.CreateQuestion)
			auth.POST("/question/:id/createAnswer", controllers.CreateAnswer)
			auth.POST("/question/:id/vote", controllers.VoteHandler("question"))
			auth.POST("/answer/:id/vote", controllers.VoteHandler("answer"))
			auth.GET("/question/:id/summary", controllers.SummarizeQuestion)
			auth.GET("/getQnA/:id", controllers.GetQuestionWithAnswers)
		}
	}
}
