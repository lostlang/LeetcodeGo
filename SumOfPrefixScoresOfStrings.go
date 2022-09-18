package leetcode

func sumPrefixScores(words []string) []int {
	var out = make([]int, len(words))
	var worldMap = make(map[string]int)

	for _, word := range words {
		worldMap[word] += 1
	}

	var rWords = make([][]rune, len(words))

	for i, word := range words {
		rWords[i] = []rune(word)
	}

	var prefixScores = make(map[string]int)
	var wordsPrefixes = make(map[string][]string)

	for _, word := range rWords {
		var prefix string

		for _, char := range word {
			prefix += string(char)
			prefixScores[prefix]++
			wordsPrefixes[string(word)] = append(wordsPrefixes[string(word)], prefix)
		}
	}

	for i, word := range words {
		for _, prefix := range wordsPrefixes[word] {
			out[i] += prefixScores[prefix]
		}
	}

	for i, word := range words {
		out[i] /= worldMap[word]
	}

	return out
}
