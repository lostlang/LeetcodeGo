package leetcode

import "strings"

func minPartitions(n string) int {
	switch {
	case strings.ContainsRune(n, '9'):
		return 9
	case strings.ContainsRune(n, '8'):
		return 8
	case strings.ContainsRune(n, '7'):
		return 7
	case strings.ContainsRune(n, '6'):
		return 6
	case strings.ContainsRune(n, '5'):
		return 5
	case strings.ContainsRune(n, '4'):
		return 4
	case strings.ContainsRune(n, '3'):
		return 3
	case strings.ContainsRune(n, '2'):
		return 2
	default:
		return 1
	}
}
