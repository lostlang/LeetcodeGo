package leetcode

func countConsistentStrings(allowed string, words []string) int {
	output := 0
	mapAllowed := make(map[rune]bool)
	for _, char := range allowed {
		mapAllowed[char] = true
	}

	for _, word := range words {
		isConsistent := true
		for _, char := range word {
			if !mapAllowed[char] {
				isConsistent = false
				break
			}
		}

		if isConsistent {
			output++
		}
	}

	return output
}
