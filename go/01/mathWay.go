// https://www.hackerrank.com/contests/projecteuler/challenges/euler001/problem
package main

import (
	"fmt"
)

func calcSum(x, y int) int {
	return x * y * (y + 1) / 2
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		// find out how many numbers divisible by ...
		z := n - 1
		a := z / 3
		b := z / 5
		c := z / 15

		// calculate the sum
		sum := calcSum(3, a) + calcSum(5, b) - calcSum(15, c)

		fmt.Println(sum)
	}
}
