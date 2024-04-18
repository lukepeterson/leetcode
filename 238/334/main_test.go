package main

import "testing"

func Test_increasingTriplet(t *testing.T) {
	result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	if result != true {
		t.Error("Expected `true`, but got ", result)
	}
}
