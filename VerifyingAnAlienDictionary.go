package leetcode

func isAlienSorted(words []string, order string) bool {
	var hash = make(map[byte]int)

	for i := range order {
		hash[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if comparatorStringsByHash(words[i-1], words[i], hash) {
			return false
		}
	}

	return true
}

func comparatorStringsByHash(s1, s2 string, hash map[byte]int) bool {
	for i := 0; i < len(s1); i++ {
		if i == len(s2) {
			return true
		}

		if hash[s1[i]] == hash[s2[i]] {
			continue
		}

		if hash[s1[i]] < hash[s2[i]] {
			return false
		} else {
			return true
		}

	}

	return false
}
