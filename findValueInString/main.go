package main

import "strings"

func finalValueAfterOperations(operations []string) int {
	ans := 0

	for i := 0; i < len(operations); i++ {
		text := operations[i]

		if strings.Contains(text, "+") {
			ans = ans + 1
		} else {
			ans = ans - 1
		}
	}

	return ans
}
