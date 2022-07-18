package leetcode

import "strings"

func finalValueAfterOperations(operations []string) int {
	var out int

	for _, string := range operations {
		if strings.Contains(string, "+") {
			out++
		} else {
			out--
		}
	}

	return out
}
