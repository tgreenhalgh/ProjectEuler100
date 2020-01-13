package main

import "fmt"

// find the nth prime number
func nthPrime(n int) int {
	count := 0
	number := 1
	for count < n {
		number++
		if isPrime(number) {
			count++
		}
	}
	return number
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(nthPrime(5))
}
