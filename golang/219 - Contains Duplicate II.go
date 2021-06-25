func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = true
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}