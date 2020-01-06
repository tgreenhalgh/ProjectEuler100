package main

import (
	"testing"
)

const (
	TestFunc      = "FiboEvenSum"
	TestFuncColor = "(\033[1;36m%v\033[0m) = "
	TestGotColor  = "\033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestFuncColor + TestGotColor + TestWantColor
)

func TestFiboEvenSum(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{10, 44},
		{18, 3382},
		{23, 60696},
		{43, 350704366},
	}

	for _, test := range tests {
		if got := FiboEvenSum(test.input); got != test.want {
			t.Errorf(ColorPrint, test.input, got, test.want)
		}
	}
}
