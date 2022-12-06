package leetcode

func longestPalindromeSub(s string) string {
	var out string

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if len(s[i:j]) > len(out) &&
				isPalindromeString(s[i:j]) {
				out = s[i:j]
			}
		}
	}

	return out
}

func isPalindromeString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
