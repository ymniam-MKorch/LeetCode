func maximumWealth(accounts [][]int) int {
	var maxWealth int
	for _, p := range accounts {
		var pw int
		for _, w := range p {
			pw += w
		}
		maxWealth = max(maxWealth, pw)
	}
	return maxWealth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}