package leetcode

func findTheDifference(s string, t string) byte {
	var hashS = stringToHashmapByte(s)
	var hashT = stringToHashmapByte(t)

	for char := range hashT {
		if hashS[char] != hashT[char] {
			return char
		}
	}

	return 0
}
