package leetcode

import (
	"strconv"
)

func numDecodings(s string) int {
	return numDecodingsDP(s, make(map[string]int))
}

func numDecodingsDP(s string, hash map[string]int) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	var flag = 1

	if len(s) >= 2 {
		var conv, _ = strconv.Atoi(s[0:2])
		if conv <= 26 {
			if conv == 10 || conv == 20 {
				flag++
			}
			flag++
		}

		if conv > 26 && conv%10 == 0 {
			flag--
		}
	}

	if len(s) == 2 {
		if flag == 3 {
			flag = 1
		}
		return flag
	}

	var out = 0

	if flag < 3 {
		if val, ok := hash[s[1:]]; ok {
			out = val
		} else {
			out = numDecodingsDP(s[1:], hash)
			hash[s[1:]] = out
		}

		if flag == 2 {
			if val, ok := hash[s[2:]]; ok {
				out += val
			} else {
				shifting := numDecodingsDP(s[2:], hash)
				out += shifting
				hash[s[2:]] = shifting
			}
		}
	} else {
		if val, ok := hash[s[2:]]; ok {
			out = val
		} else {
			out = numDecodingsDP(s[2:], hash)
			hash[s[2:]] = out
		}
	}

	return out
}
