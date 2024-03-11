package slidingwindowcounter

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowCounter struct {
	WindowDuration time.Duration
	Limit          int
	Counts         map[time.Time]int
	mu             sync.Mutex
}

func NewSlidingWindowCounter(windowDuration time.Duration, limit int) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		WindowDuration: windowDuration,
		Limit:          limit,
		Counts:         make(map[time.Time]int),
	}
}

func (c *SlidingWindowCounter) Increment() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	// Remove counts older than the window
	for timestamp := range c.Counts {
		if timestamp.Before(now.Add(-c.WindowDuration)) {
			delete(c.Counts, timestamp)
		}
	}

	// Increment count for the current time
	c.Counts[now]++

	// Check if count exceeds the limit
	totalCount := 0
	for _, count := range c.Counts {
		totalCount += count
	}
	if totalCount > c.Limit {
		// Rollback the increment
		c.Counts[now]--
		return false
	}
	return true
}

func main() {
	counter := NewSlidingWindowCounter(5*time.Second, 3) // 3 requests allowed in 5 seconds

	// Simulate some requests
	for i := 0; i < 10; i++ {
		if counter.Increment() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied (Rate limit exceeded)\n", i+1)
		}
		time.Sleep(time.Second) // Simulate some delay between requests
	}
}
