package main

import (
	"testing"
)

const (
	TestFunc      = "smallestMult"
	TestFuncColor = "(\033[1;36m%v\033[0m) = "
	TestGotColor  = "\033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestFuncColor + TestGotColor + TestWantColor
)

func TestSmallestMult(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{5, 60},
		{7, 420},
		{10, 2520},
		{13, 360360},
		{20, 232792560},
	}

	for _, test := range tests {
		if got := smallestMult(test.input); got != test.want {
			t.Errorf(ColorPrint, test.input, got, test.want)
		}
	}
}
