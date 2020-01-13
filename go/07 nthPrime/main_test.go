package main

import (
	"testing"
)

const (
	TestFunc      = "nthPrime"
	TestFuncColor = "(\033[1;36m%v\033[0m) = "
	TestGotColor  = "\033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestFuncColor + TestGotColor + TestWantColor
)

func TestNthPrime(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{6, 13},
		{10, 29},
		{100, 541},
		{1000, 7919},
		{10001, 104743},
	}

	for _, test := range tests {
		if got := nthPrime(test.input); got != test.want {
			t.Errorf(ColorPrint, test.input, got, test.want)
		}
	}
}
