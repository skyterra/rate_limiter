package rate_limiter

import (
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/time/rate"
)

// Package rate_limiter provides in-memory rate limiter management
// using LRU cache to store rate limiter instances per key.
//
// Features:
// 1. LRU cache: caches up to 10240 limiters; least recently used are evicted.
// 2. Flexible creation: NewRateLimiter allows specifying requests per second and burst.
// 3. Safe access: GetRateLimiter returns nil if limiter does not exist.
// 4. Use cases:
//    - API request limiting
//    - Anti-bot operations
//    - Internal service call rate control
//
// Note:
// - This is a single-node in-memory limiter, not suitable for distributed scenarios.
// - Configure parameters carefully according to your application's needs.

var rc *lru.Cache

func init() {
	rc, _ = lru.New(10240)
}

// GetRateLimiter retrieves an existing rate limiter by key.
// Returns nil if no limiter exists.
func GetRateLimiter(key string) *rate.Limiter {
	v, ok := rc.Get(key)
	if !ok {
		return nil
	}
	limiter, _ := v.(*rate.Limiter)
	return limiter
}

// NewRateLimiter creates a new rate limiter and caches it in the global LRU.
// limit: maximum operations per second
// burst: maximum burst operations allowed
func NewRateLimiter(key string, limit float64, burst int) *rate.Limiter {
	limiter := rate.NewLimiter(rate.Limit(limit), burst)
	rc.Add(key, limiter)
	return limiter
}
