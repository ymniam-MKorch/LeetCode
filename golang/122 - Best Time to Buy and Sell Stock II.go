func maxProfit(prices []int) int {
	var res int
	if len(prices) == 0 {
		return res
	}
	start := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			res = res + prices[i-1] - prices[start]
			start = i
		}
	}
	if start != len(prices)-1 {
		res = res + prices[len(prices)-1] - prices[start]
	}
	return res
}