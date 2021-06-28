func distributeCandies(candyType []int) int {
	n := len(candyType)
	m := make(map[int]bool)
	for _, v := range candyType {
		m[v] = true
	}
	return min(len(m), n/2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}