package leakybucket

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	Capacity     int
	LeakInterval time.Duration
	LeakAmount   int
	LastLeakTime time.Time
	CurrentLevel int
	mu           sync.Mutex
}

func NewLeakyBucket(capacity, leakAmount int, leakInterval time.Duration) *LeakyBucket {
	return &LeakyBucket{
		Capacity:     capacity,
		LeakAmount:   leakAmount,
		LeakInterval: leakInterval,
		LastLeakTime: time.Now(),
		CurrentLevel: 0,
	}
}

func (b *LeakyBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Calculate elapsed time since last leak
	currentTime := time.Now()
	elapsed := currentTime.Sub(b.LastLeakTime)

	// Leak water from the bucket
	leakAmount := int(elapsed / b.LeakInterval) * b.LeakAmount
	if leakAmount > b.CurrentLevel {
		b.CurrentLevel = 0
	} else {
		b.CurrentLevel -= leakAmount
	}

	// If the bucket overflows, deny the request
	if b.CurrentLevel >= b.Capacity {
		return false
	}

	// Add water to the bucket
	b.CurrentLevel++
	b.LastLeakTime = currentTime
	return true
}

func main() {
	bucket := NewLeakyBucket(10, 1, time.Second) // Capacity: 10, LeakAmount: 1 per second

	// Simulate some requests
	for i := 0; i < 15; i++ {
		if bucket.Allow() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied (Rate limit exceeded)\n", i+1)
		}
		time.Sleep(time.Millisecond * 200) // Simulate some delay between requests
	}
}
