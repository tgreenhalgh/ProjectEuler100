/*
A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two n-digit numbers.
*/
package main

import (
	"fmt"
	"math"
)

func isPalindrome(number int) bool {
	var remainder, reverse int
	temp := number

	for temp != 0 {
		remainder = temp % 10
		reverse = reverse*10 + remainder
		temp /= 10
	}
	return number == reverse
}

func largestPalindromeProduct(x int) int {
	max := int(math.Pow(10, float64(x))) - 1
	maxProd := 0

	for i := max; i > max/2; i-- {
		for j := max; j > max/2; j-- {
			if isPalindrome(i * j) {
				if i*j > maxProd {
					maxProd = i * j
				}
			}
		}
	}
	return maxProd
}

func main() {
	fmt.Println(largestPalindromeProduct(3))
}
