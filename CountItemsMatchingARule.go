package leetcode

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var out int

	for _, item := range items {
		switch ruleKey {
		case "color":
			if item[1] == ruleValue {
				out++
			}
		case "type":
			if item[0] == ruleValue {
				out++
			}
		case "name":
			if item[2] == ruleValue {
				out++
			}
		}
	}

	return out
}
