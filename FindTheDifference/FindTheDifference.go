package FindTheDifference

import "leetcode/utils"

var StringToHashmapByte = utils.StringToHashmapByte

func findTheDifference(s string, t string) byte {
	hashS := StringToHashmapByte(s)
	hashT := StringToHashmapByte(t)

	for char := range hashT {
		if hashS[char] != hashT[char] {
			return char
		}
	}

	return 0
}
