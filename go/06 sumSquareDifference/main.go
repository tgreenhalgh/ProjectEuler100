package main

import "fmt"

func sumSquareDifference(n int) int {
	sum := 0
	sqSum := 0

	for i := 1; i <= n; i++ {
		sum += i
		sqSum += (i * i)
	}
	return sum*sum - sqSum
}

func main() {
	fmt.Println(sumSquareDifference(5))
}
