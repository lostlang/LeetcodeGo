package FindCommonCharacters

import "leetcode/utils"

var StringToHashmapRune = utils.StringToHashmapRune

func commonChars(words []string) []string {
	var out []string
	hash := StringToHashmapRune(words[0])

	for _, word := range words {
		newHash := StringToHashmapRune(word)
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
