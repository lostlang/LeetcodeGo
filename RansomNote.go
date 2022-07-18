package leetcode

func canConstruct(ransomNote string, magazine string) bool {

	var noteHash = make(map[rune]int)
	var magazineHash = make(map[rune]int)

	for _, char := range ransomNote {
		noteHash[char]++
	}

	for _, char := range magazine {
		magazineHash[char]++
	}

	for val := range noteHash {
		if noteHash[val] > magazineHash[val] {
			return false
		}
	}
	return true
}
