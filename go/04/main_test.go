package main

import (
	"testing"
)

const (
	TestFunc      = "largestPalindromeProduct"
	TestFuncColor = "(\033[1;36m%v\033[0m) = "
	TestGotColor  = "\033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestFuncColor + TestGotColor + TestWantColor
)

func TestLargestPalindromeProduct(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 9009},
		{3, 906609},
	}

	for _, test := range tests {
		if got := largestPalindromeProduct(test.input); got != test.want {
			t.Errorf(ColorPrint, test.input, got, test.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(b.N)
	}
}
