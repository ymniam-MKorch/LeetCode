func maxProfit(prices []int) int {
	var start, res int
	if len(prices) == 0 {
		return 0
	}
	start = prices[0]
	for _, v := range prices {
		res = max(res, v-start)
		if v < start {
			start = v
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}