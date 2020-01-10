package main

import (
	"testing"
)

const (
	TestFunc      = "largestPrimeFactor"
	TestFuncColor = "(\033[1;36m%v\033[0m) = "
	TestGotColor  = "\033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestFuncColor + TestGotColor + TestWantColor
)

func TestLargestPrimeFactor(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 2},
		{3, 3},
		{5, 5},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, test := range tests {
		if got := largestPrimeFactor(test.input); got != test.want {
			t.Errorf(ColorPrint, test.input, got, test.want)
		}
	}
}
