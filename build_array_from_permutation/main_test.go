package main

import "testing"

func TestBuildArray(t *testing.T) {
	given := []int{0, 2, 1, 5, 3, 4}
	expected := []int{0, 1, 2, 4, 5, 3}

	get := buildArray(given)

	if !isEqual(expected, get) {
		t.Errorf("expected not equals get")
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
