package leetcode

func partitionString(s string) int {
	var out int

	var hash = make(map[rune]int)

	for _, r := range s {
		if hash[r] == 0 {
			hash[r]++
		} else {
			out++
			hash = make(map[rune]int)
			hash[r]++
		}
	}

	for _, v := range hash {
		if v == 1 {
			out++
			break
		}
	}

	return out
}
