package leetcode

func repeatedCharacter(s string) byte {
	var hash = make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if hash[s[i]] {
			return s[i]
		} else {
			hash[s[i]] = true
		}
	}

	return 0
}
