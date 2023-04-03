package leetcode

func checkIfPangram(sentence string) bool {
	output := false
	letters := make(map[rune]bool)
	for _, letter := range sentence {
		letters[letter] = true
	}
	if len(letters) == 26 {
		output = true
	}

	return output
}
