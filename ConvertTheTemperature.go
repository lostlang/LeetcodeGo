package leetcode

func convertTemperature(celsius float64) []float64 {
	output := make([]float64, 2)
	output[0] = celsius + 273.15
	output[1] = celsius*1.8 + 32

	return output
}
