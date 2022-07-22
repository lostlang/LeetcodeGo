package leetcode

func longestPalindrome(s string) int {
	var out int
	var odd bool

	var hashS = make(map[rune]int)

	for _, char := range s {
		hashS[char] += 1
	}

	for char := range hashS {
		out += hashS[char] / 2 * 2
		if hashS[char]%2 == 1 {
			odd = true
		}
	}

	if odd {
		out++
	}

	return out
}
