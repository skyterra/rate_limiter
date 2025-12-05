package rate_limiter

import (
	"testing"
	"time"
)

func TestRateLimiterLifecycle(t *testing.T) {
	key := "user:123"
	limit := 10.0
	burst := 20

	l := NewRateLimiter(key, limit, burst)
	if l == nil {
		t.Fatal("Failed to create limiter")
	}

	got := GetRateLimiter(key)
	if got == nil {
		t.Fatal("Failed to get existing limiter")
	}

	if got != l {
		t.Error("Got pointer does not match created pointer")
	}

	if GetRateLimiter("non-existent") != nil {
		t.Error("Should return nil for non-existent key")
	}
}

func TestRateLimitingLogic(t *testing.T) {
	key := "api:fast"
	l := NewRateLimiter(key, 1, 1)

	if !l.Allow() {
		t.Error("First request should be allowed")
	}

	if l.Allow() {
		t.Error("Second request should be rejected immediately")
	}

	time.Sleep(1100 * time.Millisecond)
	if !l.Allow() {
		t.Error("Request should be allowed after waiting")
	}
}
