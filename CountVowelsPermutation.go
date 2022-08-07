package leetcode

func countVowelPermutation(n int) int {
	var out int
	var hash = map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}

	var next map[byte]int

	for i := 1; i < n; i++ {
		next = make(map[byte]int)
		for k, v := range hash {
			switch k {
			case 'a':
				next['e'] += v
			case 'e':
				next['i'] += v
				next['a'] += v
			case 'i':
				next['a'] += v
				next['e'] += v
				next['o'] += v
				next['u'] += v
			case 'o':
				next['i'] += v
				next['u'] += v
			case 'u':
				next['a'] += v
			}
		}

		for k, v := range next {
			hash[k] = v % 1000000007
		}
	}

	for _, v := range hash {
		out += v
	}

	out %= 1000000007

	return out
}
