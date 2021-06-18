package util

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
