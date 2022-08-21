package leetcode

func largestPalindromic(num string) string {
	var out string
	var hash = make(map[rune]int)

	for _, r := range num {
		hash[r]++
	}

	var left, right, center string

	for i := '1'; i <= '9'; i++ {
		for hash[i] > 1 {
			left += string(i)
			right = string(i) + right
			hash[i] -= 2
		}
	}

	if len(left) > 0 {
		for hash['0'] > 1 {
			left = "0" + left
			right += "0"
			hash['0'] -= 2
		}
	}

	for i := '9'; i >= '0'; i-- {
		if hash[i] > 0 {
			center += string(i)
			hash[i]--
			break
		}
	}

	out = right + center + left

	return out
}
