// cmd/api/main.go
package main

import (
	"github.com/ashblend17/stackoverflow-sample/config"
	"github.com/ashblend17/stackoverflow-sample/database"
	"github.com/ashblend17/stackoverflow-sample/middlewares"
	"github.com/ashblend17/stackoverflow-sample/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	r := gin.Default()
	r.Use(middlewares.GlobalUserOrIPRateLimitMiddleware())

	routes.RegisterRoutes(r)

	r.Run(":8080") // starts server
}
