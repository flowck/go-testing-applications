package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4

	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', expected: '%v'", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 2 - 2
	expected := 3

	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', expected: '%v'", got, expected)
	}
}
