package CellsInARangeOnAnExcelSheet

func cellsInRange(s string) []string {
	var out []string
	rS := []rune(s)
	var start, end [2]rune

	start[0] = rS[0]
	start[1] = rS[1]
	end[0] = rS[3]
	end[1] = rS[4]

	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			out = append(out, string([]rune{rune(i), rune(j)}))
		}
	}

	return out
}
