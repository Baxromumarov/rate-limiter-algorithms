package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity       int
	Tokens         int
	RefillRate     int
	LastRefillTime time.Time
	mu             sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		Capacity:       capacity,
		Tokens:         capacity,
		RefillRate:     refillRate,
		LastRefillTime: time.Now(),
	}
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	timeSinceLastRefill := now.Sub(b.LastRefillTime)

	// Calculate how many tokens should be added since last refill
	tokensToAdd := int(timeSinceLastRefill.Seconds()) * b.RefillRate

	// Add tokens up to the capacity
	b.Tokens = min(b.Tokens+tokensToAdd, b.Capacity)

	// Consume token if available
	if b.Tokens > 0 {
		b.Tokens--
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	bucket := NewTokenBucket(10, 1) // Capacity: 10, RefillRate: 1 token per second

	// Simulate some requests
	for i := 0; i < 30; i++ {
		if bucket.Allow() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied (Rate limit exceeded)\n", i+1)
		}
		time.Sleep(time.Millisecond * 20) // Simulate some delay between requests
	}
}
