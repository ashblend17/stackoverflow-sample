package middlewares

import (
	"net/http"
	"strings"

	"github.com/ashblend17/stackoverflow-sample/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			ctx.Abort()
			return
		}

		// Extract token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify and get user ID
		userID, err := utils.VerifyJWT(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Attach user ID to context
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
