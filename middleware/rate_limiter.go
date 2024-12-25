package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Visitor struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}

type RateLimiter struct {
	Visitors map[string]*Visitor
	Mu       sync.Mutex
	Rate     rate.Limit
	Burst    int
}

func NewRateLimiter(rps int, burst int) *RateLimiter {
	rl := &RateLimiter{
		Visitors: make(map[string]*Visitor),
		Rate:     rate.Limit(rps),
		Burst:    burst,
	}
	go rl.cleanupVisitors()
	return rl
}

func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		rl.Mu.Lock()
		for ip, visitor := range rl.Visitors {
			if time.Since(visitor.LastSeen) > 3*time.Minute {
				delete(rl.Visitors, ip)
			}
		}
		rl.Mu.Unlock()
	}
}

func (rl *RateLimiter) getVisitor(key string) *Visitor {
	rl.Mu.Lock()
	defer rl.Mu.Unlock()
	v, exists := rl.Visitors[key]
	if !exists {
		limiter := rate.NewLimiter(rl.Rate, rl.Burst)
		v = &Visitor{Limiter: limiter, LastSeen: time.Now()}
		rl.Visitors[key] = v
	}
	v.LastSeen = time.Now()
	return v
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		visitor := rl.getVisitor(key)

		if !visitor.Limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try again later.",
			})
			return
		}

		c.Next()
	}
}
