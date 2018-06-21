package main

import "testing"

func TestReturnTrueForFirstTest(t *testing.T) {
	result := returnTrueForFirstTest()
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: %t.", result, true)
	}
}
