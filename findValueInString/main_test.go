package main

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	given := []string{"X++", "++X", "--X", "X--"}
	want := 0

	get := finalValueAfterOperations(given)

	if get != want {
		t.Errorf("Failed get %v, want %v", get, want)
	}
}
