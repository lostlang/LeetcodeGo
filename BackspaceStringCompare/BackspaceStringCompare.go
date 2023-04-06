package BackspaceStringCompare

func backspaceCompare(s string, t string) bool {
	s = backspaceConvert(s)
	t = backspaceConvert(t)

	return s == t
}

func backspaceConvert(s string) string {
	rS := []rune(s)

	i := 0

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
