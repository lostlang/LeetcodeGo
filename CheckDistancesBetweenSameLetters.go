package leetcode

func checkDistances(s string, distance []int) bool {
	var rS = []rune(s)

	for i, char := range rS {
		if char >= 'a' {
			var next = i + distance[int(char-'a')] + 1
			if next >= len(rS) || char != rS[next] {
				return false
			} else {
				rS[next] = 0
			}
		}
	}

	return true
}
