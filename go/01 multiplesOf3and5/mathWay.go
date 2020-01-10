// https://www.hackerrank.com/contests/projecteuler/challenges/euler001/problem
package main

func calcSum(x, y int) int {
	return x * y * (y + 1) / 2
}

// Fancy calculates sum quickly
func Fancy(n int) int {
	z := n - 1
	a := z / 3
	b := z / 5
	c := z / 15

	// calculate the sum
	sum := calcSum(3, a) + calcSum(5, b) - calcSum(15, c)

	return sum
}
