package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limiterStore = make(map[string]*rate.Limiter)
	mu           sync.Mutex

	// Rate: 1 requests per second, with a burst of 5
	rateLimit     = rate.Limit(1)
	burstCapacity = 1
)

func getLimiter(key string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, exists := limiterStore[key]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(rateLimit, burstCapacity)
	limiterStore[key] = limiter
	return limiter
}

func GlobalUserOrIPRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string

		// Try to use user ID (from AuthMiddleware)
		if uid, exists := c.Get("user_id"); exists {
			key = "user:" + string(rune(uid.(int))) // basic conversion
		} else {
			// Fallback to IP
			key = "ip:" + c.ClientIP()
		}

		limiter := getLimiter(key)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests (user/ip rate limit)",
			})
			return
		}

		c.Next()
	}
}
