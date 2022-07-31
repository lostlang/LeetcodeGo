package leetcode

func backspaceCompare(s string, t string) bool {
	s = backspaceConvert(s)
	t = backspaceConvert(t)

	return s == t
}

func backspaceConvert(s string) string {
	var rS = []rune(s)

	var i int

	for i < len(rS) {
		if rS[i] != '#' {
			i++
		} else {
			if i > 0 {
				rS = append(rS[:i-1], rS[i+1:]...)
				i--
			} else {
				rS = rS[1:]
			}
		}

	}

	return string(rS)
}
