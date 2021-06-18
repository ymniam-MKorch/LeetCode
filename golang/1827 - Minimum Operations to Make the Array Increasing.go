func minOperations(nums []int) int {
	ans := 0
	if len(nums) < 2 {
		return ans
	}
	for i := 1; i < len(nums); i++ {
		for nums[i] <= nums[i-1] {
			nums[i]++
			ans++
		}
	}
	return ans
}