package leetcode

func restoreString(s string, indices []int) string {
	var out = make([]byte, len(s))

	for i, v := range indices {
		out[v] = s[i]
	}

	return string(out)
}
