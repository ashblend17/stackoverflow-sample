package utils

import (
	"sync"

	"golang.org/x/time/rate"
)

var (
	voteLimiters   = make(map[int]*rate.Limiter) // user_id → limiter
	voteLimiterMux sync.Mutex
)

// NewLimiter creates a rate limiter for 1 req/sec, burst of 5
func GetUserVoteLimiter(userID int) *rate.Limiter {
	voteLimiterMux.Lock()
	defer voteLimiterMux.Unlock()

	limiter, exists := voteLimiters[userID]
	if !exists {
		limiter = rate.NewLimiter(1, 5) // 1 vote/sec, burst 5
		voteLimiters[userID] = limiter
	}
	return limiter
}
