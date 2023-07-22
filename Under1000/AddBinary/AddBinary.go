package AddBinary

func addBinary(a string, b string) string {
	output := ""
	tmpBit := false

	if len(a) < len(b) {
		a, b = b, a
	}

	aR := []rune(a)
	bR := []rune(b)
	dif := len(aR) - len(bR)

	for i := len(aR) - 1; i >= 0; i-- {
		if i >= dif {
			if aR[i] == bR[i-dif] {
				if tmpBit {
					output = "1" + output
				} else {
					output = "0" + output
				}

				if aR[i] == '1' {
					tmpBit = true
				} else {
					tmpBit = false
				}
			} else {
				if !tmpBit {
					output = "1" + output
				} else {
					output = "0" + output
				}
			}
		} else {
			if tmpBit {
				if aR[i] == '1' {
					output = "0" + output
				} else {
					output = "1" + output
					tmpBit = false
				}
			} else {
				output = string(aR[i]) + output
			}
		}
	}

	if tmpBit {
		output = "1" + output
	}

	return output
}
