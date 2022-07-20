package leetcode

func numMatchingSubseq(s string, words []string) int {
	var out int
	var sumLen = make([]int, len(words))

	for ind := range s {

		for ind2 := range words {
			if sumLen[ind2] == len(words[ind2])+1 {
				continue
			}

			if s[ind] == words[ind2][sumLen[ind2]] {
				sumLen[ind2]++
			}

			if sumLen[ind2] == len(words[ind2]) {
				out++
				sumLen[ind2]++
			}

		}
	}

	return out
}
