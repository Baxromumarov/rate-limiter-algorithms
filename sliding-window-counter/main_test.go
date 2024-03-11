package slidingwindowcounter_test

import (
	swc "baxromumarov/rate-limiter-algorithms/sliding-window-counter"
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	// Maximum 10 requests per minute
	counter := swc.NewSlidingWindowCounter(time.Second*15, 10)

	// Simulate some requests
	for i := 0; i < 15; i++ {
		allowed := counter.Increment()
		fmt.Printf("Request: %v, result of algorithm: %v", i, allowed)
		if (i == 11 || i == 12 || i == 13 || i == 14 || i == 15) && allowed != false {
			fmt.Printf("i: %d, result: %t", i, allowed)
			t.Fail()
		} else {
			fmt.Printf("i: %d, result: %t", i, allowed)
		}

	}

}
