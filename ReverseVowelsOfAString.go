package leetcode

func reverseVowels(s string) string {
	var swapString []rune
	var outString = []rune(s)

	for _, char := range outString {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			swapString = append([]rune{char}, swapString...)
		}
	}

	for i, char := range outString {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			outString[i] = swapString[0]
			swapString = swapString[1:]
		}
	}

	return string(outString)
}
