package main

import (
    "fmt"
    "sync"
    "time"
)

type FixedWindowCounter struct {
    windowSize   time.Duration
    maxRequests  int
    timestamps   []time.Time
    mu           sync.Mutex
}

func NewFixedWindowCounter(windowSize time.Duration, maxRequests int) *FixedWindowCounter {
    return &FixedWindowCounter{
        windowSize:  windowSize,
        maxRequests: maxRequests,
        timestamps:  make([]time.Time, 0),
    }
}

func (c *FixedWindowCounter) Increment() bool {
    c.mu.Lock()
    defer c.mu.Unlock()

    currentTime := time.Now()
    c.timestamps = append(c.timestamps, currentTime)

    // Remove timestamps older than the window
    var validTimestamps []time.Time
    for _, ts := range c.timestamps {
        if currentTime.Sub(ts) <= c.windowSize {
            validTimestamps = append(validTimestamps, ts)
        }
    }
    c.timestamps = validTimestamps

    // Check if number of requests exceeds the limit
    if len(c.timestamps) > c.maxRequests {
        return false
    }
    return true
}

func main() {
    windowSize := 60 * time.Second   // 60 seconds
    maxRequests := 10                 // Maximum 10 requests per minute
    counter := NewFixedWindowCounter(windowSize, maxRequests)

    // Simulate some requests
    for i := 0; i < 15; i++ {
        allowed := counter.Increment()
        if allowed {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Denied (Rate limit exceeded)\n", i+1)
        }
    }
}
