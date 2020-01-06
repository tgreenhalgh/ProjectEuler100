package main

import "fmt"

// FiboEvenSum returns the sum of the even terms in the Fibonacci sequence
// that do not exceed the nth term
func FiboEvenSum(n int) int {
	var sum int
	fibs := []int{1, 2}

	for i := 0; i < n-2; i++ {
		fibs = append(fibs, fibs[i]+fibs[i+1])
	}

	for _, e := range fibs {
		if e%2 == 0 {
			sum += e
		}
	}

	return sum
}

func main() {
	fmt.Println(FiboEvenSum(10))
}
