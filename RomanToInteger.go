package leetcode

func romanToInt(s string) int {
	var alphabet = make(map[rune]int)
	alphabet['I'] = 1
	alphabet['V'] = 5
	alphabet['X'] = 10
	alphabet['L'] = 50
	alphabet['C'] = 100
	alphabet['D'] = 500
	alphabet['M'] = 1000

	var sum = 0
	var curSum = 0
	var prev = 0

	for _, char := range s {
		if alphabet[char] < prev {
			sum += curSum
			curSum = 0
			prev = alphabet[char]
		}

		if alphabet[char] > curSum && curSum != 0 {
			sum += alphabet[char] - curSum
			curSum = 0
			prev = alphabet[char]
		} else {
			prev = alphabet[char]
			curSum += alphabet[char]
		}
	}
	sum += curSum

	return sum
}
