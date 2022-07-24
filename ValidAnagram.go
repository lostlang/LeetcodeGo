package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var hashS = make(map[rune]int)
	var hashT = make(map[rune]int)

	for _, char := range s {
		hashS[char]++
	}

	for _, char := range t {
		hashT[char]++
	}

	for i := range hashS {
		if hashS[i] != hashT[i] {
			return false
		}
	}

	return true
}
