package leetcode

func longestCommonPrefix(strs []string) string {
	var out string
	var minLenWord = len(strs[0])
	var indexMinWord = 0
	var charArray [][]rune = make([][]rune, len(strs))

	for index, str := range strs {
		if minLenWord > len(str) {
			minLenWord = len(str)
			indexMinWord = index
		}
		charArray[index] = []rune(str)
	}

	for index, char := range strs[indexMinWord] {
		for _, str := range charArray {
			if str[index] != char {
				return out
			}
		}
		out = out + string(char)
	}

	return out
}
