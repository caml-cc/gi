package server

import (
	"net/http"
	"sync"
	"time"
)

var (
	requestCount int
	windowStart  time.Time
	mu           sync.Mutex
)

func rateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()

		if now.Sub(windowStart) >= time.Minute {
			requestCount = 0
			windowStart = now
		}

		if requestCount >= 1000 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		requestCount++
		next.ServeHTTP(w, r)
	})
}
