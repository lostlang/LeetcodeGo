package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Runes = []rune(s1)
	var s2Runes = []rune(s2)
	start := 0

	for start+len(s1Runes) <= len(s2Runes) {
		if containSubString(s1Runes,
			s2Runes[start:start+len(s1Runes)]) {
			return true
		} else {
			start++
		}
	}

	return false
}

func containSubString(s1 []rune, s2 []rune) bool {
	var s1Map = make(map[rune]int)
	var s2Map = make(map[rune]int)

	for _, v := range s1 {
		s1Map[v]++
	}
	for _, v := range s2 {
		s2Map[v]++
	}

	for k, v := range s1Map {
		if s2Map[k] != v {
			return false
		}
	}

	return true
}
