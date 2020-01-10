package main

import (
	"fmt"
	"testing"
)

/*
multiplesOf3and5(1000) should return 233168.
multiplesOf3and5(49) should return 543.
multiplesOf3and5(19564) should return 89301183.
multiplesOf3and5(8456) should return 16687353.
*/
func TestMultiplesOf3and5(t *testing.T) {
	fmt.Println("Running tests...")
	results := MultiplesOf3and5(1000)
	expected := 233168
	if results != expected {
		t.Errorf("Expected %d got %d", expected, results)
	}

	results = MultiplesOf3and5(49)
	expected = 543
	if results != expected {
		t.Errorf("Expected %d got %d", expected, results)
	}

	results = MultiplesOf3and5(19564)
	expected = 89301183
	if results != expected {
		t.Errorf("Expected %d got %d", expected, results)
	}

	results = MultiplesOf3and5(8456)
	expected = 16687353
	if results != expected {
		t.Errorf("Expected %d got %d", expected, results)
	}
}

func TestFancy(t *testing.T) {
	fmt.Println("Testing fancy")
	results := Fancy(1000)
	expected := 233168
	if results != expected {
		t.Errorf("Expected %d got %d", expected, results)
	}
}
