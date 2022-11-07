package main

import "testing"

func TestRunningSum(t *testing.T) {
	given := []int{1, 2, 3, 4}
	expect := []int{1, 3, 6, 10}

	get := runningSum(given)

	if !isEqual(expect, get) {
		t.Errorf("Failed")
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
