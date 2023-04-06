package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)
	output := 0

	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			output++
			i++
			j++
		} else {
			j++
		}
	}

	return output
}
