package leetcode

func isAlienSorted(words []string, order string) bool {
	// var hash = make(map[byte]int)
	//
	// for i := range order {
	// 	hash[order[i]] = i
	// }
	//
	// for i := 1; i < len(words); i++ {
	// 	if len(words[i]) > len(words[i-1]) {
	// 		for j := range words[i-1] {
	// 			if hash[words[i][j]] > hash[words[i-1][j]] {
	// 				return false
	// 			}
	// 		}
	// 	} else {
	// 		for j := range words[i] {
	// 			if hash[words[i][j]] > hash[words[i-1][j]] {
	// 				return false
	// 			}
	// 		}
	// 	}
	//
	// }

	return true
}
