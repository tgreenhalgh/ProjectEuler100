package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(x int) int {
	var maxPrime int

	// parse out the evens
	for x%2 == 0 {
		maxPrime = 2
		x /= 2
	}

	// parse out the odds
	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		for x%i == 0 {
			maxPrime = i
			x /= i
		}
	}

	// x is prime at this point
	if x > 2 {
		maxPrime = x
	}

	return maxPrime
}

func main() {
	fmt.Println(largestPrimeFactor(0))
}
