package leetcode

func addBinary(a string, b string) string {
	var out string
	var tmpBit bool

	if len(a) < len(b) {
		a, b = b, a
	}

	var aR = []rune(a)
	var bR = []rune(b)
	var dif = len(aR) - len(bR)

	for i := len(aR) - 1; i >= 0; i-- {

		if i >= dif {
			if aR[i] == bR[i-dif] {
				if tmpBit {
					out = "1" + out
				} else {
					out = "0" + out
				}

				if aR[i] == '1' {
					tmpBit = true
				} else {
					tmpBit = false
				}
			} else {
				if !tmpBit {
					out = "1" + out
				} else {
					out = "0" + out
				}
			}
		} else {
			if tmpBit {
				if aR[i] == '1' {
					out = "0" + out
				} else {
					out = "1" + out
					tmpBit = false
				}
			} else {
				out = string(aR[i]) + out
			}
		}
	}

	if tmpBit {
		out = "1" + out
	}

	return out
}
