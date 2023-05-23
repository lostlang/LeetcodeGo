package utils

func StringToHashmapByte(str string) map[byte]int {
	hash := make(map[byte]int)

	for i := range str {
		hash[str[i]]++
	}

	return hash
}

func StringToHashmapRune(str string) map[rune]int {
	hash := make(map[rune]int)

	for _, val := range str {
		hash[val]++
	}

	return hash
}

func IntToHash(arr []int) map[int]int {
	out := make(map[int]int)

	for _, val := range arr {
		out[val]++
	}

	return out
}
