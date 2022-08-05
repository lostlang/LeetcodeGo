package leetcode

func commonChars(words []string) []string {
	var out []string
	var hash = stringToHashmapRune(words[0])

	for _, word := range words {
		var newHash = stringToHashmapRune(word)
		for key := range hash {
			if hash[key] > newHash[key] {
				hash[key] = newHash[key]
			}
		}
	}

	for char, count := range hash {
		for i := 0; i < count; i++ {
			out = append(out, string(char))
		}
	}

	return out
}
