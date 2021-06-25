func minMoves(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var sum int
	min := math.MaxInt32
	for _, v := range nums {
		if v < min {
			min = v
		}
		sum += v
	}
	return sum - min*len(nums)
}