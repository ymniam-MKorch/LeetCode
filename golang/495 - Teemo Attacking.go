func findPoisonedDuration(timeSeries []int, duration int) int {
	var res int
	if len(timeSeries) == 0 {
		return res
	}
	for i := 0; i < len(timeSeries)-1; i++ {
		res += min(timeSeries[i+1]-timeSeries[i], duration)
	}
	return res + duration
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}