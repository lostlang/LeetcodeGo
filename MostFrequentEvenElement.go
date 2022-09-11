package leetcode

func mostFrequentEven(nums []int) int {
	var out = -1
	var outFreq int

	var count = make(map[int]int)

	for _, num := range nums {
		if num%2 == 0 {
			count[num]++
		}
	}

	for num, freq := range count {
		if freq > outFreq || (freq == outFreq && num < out) {
			outFreq = freq
			out = num
		}
	}

	return out
}
