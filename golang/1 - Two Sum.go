func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := numMap[another]; ok {
			return []int{numMap[another], i}
		}
		numMap[nums[i]] = i
	}
	return nil
}