package main

import "fmt"

func sieveOfEratosthenes(n int) []int {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to n?
func smallestMult(n int) int {
	product := 1
	primes := sieveOfEratosthenes(n)

	// take out all the prime factors
	for _, v := range primes {
		t := 1
		for t <= n/v {
			t *= v
		}
		product *= t
		t = 1
	}

	return product
}

// 5 -> 1 * 2 * 3 * 4 * 5 = 120 -> 5 * 3 * 2^2 * 2 -> 5 * 3 * 2^2 = 60

func main() {
	fmt.Println(smallestMult(5))
}
