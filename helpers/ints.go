package helpers


func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SumOfSeries(start, end int) int {
	step := 1

	count := (end - start) / step
	sumOfExtremes := start + (end - step)

	return count * sumOfExtremes / 2
}
