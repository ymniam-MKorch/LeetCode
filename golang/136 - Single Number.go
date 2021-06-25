func singleNumber(nums []int) int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}
	return xor
}