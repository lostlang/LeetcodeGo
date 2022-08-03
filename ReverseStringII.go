package leetcode

func reverseStr(s string, k int) string {
	var reverse string
	var tmp string

	for i, char := range s {
		if i/k%2 == 0 {
			tmp = string(char) + tmp
		} else {
			reverse = reverse + tmp + string(char)
			tmp = ""

		}
	}

	reverse = reverse + tmp

	return reverse
}
