package main

import "testing"

func TestGetConcatenation(t *testing.T) {
	give := []int{1, 2, 1}
	expected := []int{1, 2, 1, 1, 2, 1}

	get := getConcatenation(give)

	if !isEqual(expected, get) {
		t.Errorf("Not eqauls expected")
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
