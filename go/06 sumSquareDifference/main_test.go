package main

import (
	"testing"
)

const (
	TestFunc      = "sumSquareDifference"
	TestFuncColor = "(\033[1;36m%v\033[0m) = "
	TestGotColor  = "\033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestFuncColor + TestGotColor + TestWantColor
)

func TestSumSquareDifference(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{10, 2640},
		{20, 41230},
		{100, 25164150},
	}

	for _, test := range tests {
		if got := sumSquareDifference(test.input); got != test.want {
			t.Errorf(ColorPrint, test.input, got, test.want)
		}
	}
}
