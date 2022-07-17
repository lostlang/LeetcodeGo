package leetcode

func rearrangeCharacters(s string, target string) int {
	var out = -1
	var stringHash = make(map[rune]int)
	var targetHash = make(map[rune]int)

	for _, char := range s {
		stringHash[char]++
	}
	for _, char := range target {
		targetHash[char]++
	}

	for char := range targetHash {
		if out == -1 {
			out = stringHash[char] / targetHash[char]
		}
		curent := stringHash[char] / targetHash[char]
		if out > curent {
			out = curent
		}

		if out == 0 {
			break
		}

	}

	return out
}
