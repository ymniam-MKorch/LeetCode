func findMaxConsecutiveOnes(nums []int) int {
	var res, len int
	for _, v := range nums {
		if v == 0 {
			res = max(res, len)
			len = 0
		} else {
			len++
		}
	}
	res = max(res, len)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}