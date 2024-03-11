package fixedwindowcounter_test

import (
	fwc "baxromumarov/rate-limiter-algorithms/fixed-window-counter"
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	windowSize := 60 * time.Second // 60 seconds
	maxRequests := 10              // Maximum 10 requests per minute
	counter := fwc.NewFixedWindowCounter(windowSize, maxRequests)

	// Simulate some requests
	for i := 0; i < 15; i++ {
		allowed := counter.Increment()
		// fmt.Printf("Request: %v, result of algorithm: %v\n", i, allowed)
	
		if (i == 10|| i == 11 || i == 12 || i == 13 || i == 14 || i == 15  ) && allowed==true {
			fmt.Printf("i: %d, result: %t\n", i, allowed)
			t.Fail()
			return
		} else {
			fmt.Printf("i: %d, result: %t\n", i, allowed)
		}

	}

}
