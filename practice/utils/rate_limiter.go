package utils

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	visitors map[string]time.Time
	mu       sync.Mutex
	limit    time.Duration
}

func NewRateLimiter(limit time.Duration) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]time.Time),
		limit:    limit,
	}
}

func (r *RateLimiter) Allow(ip string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	lastSeen, exists := r.visitors[ip]
	if !exists || time.Since(lastSeen) > r.limit {
		r.visitors[ip] = time.Now()
		return true
	}
	return false
}

func (r *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ip := req.RemoteAddr
		if !r.Allow(ip) {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, req)
	})
}
