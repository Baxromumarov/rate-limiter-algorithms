package leakybucket_test

import (
	lb "baxromumarov/rate-limiter-algorithms/leaky-bucket"
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	// Maximum 10 requests per minute
	counter := lb.NewLeakyBucket(10, 1, time.Second*2)

	// Simulate some requests
	for i := 0; i < 15; i++ {
		allowed := counter.Allow()
		fmt.Printf("Request: %v, result of algorithm: %v", i, allowed)
		if (i == 11 || i == 12 || i == 13 || i == 14 || i == 15) && allowed != false {
			fmt.Printf("i: %d, result: %t", i, allowed)
			t.Fail()
		} else {
			fmt.Printf("i: %d, result: %t", i, allowed)
		}

	}

}
