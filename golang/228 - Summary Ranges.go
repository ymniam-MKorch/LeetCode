func summaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		str := fmt.Sprintf("%d", nums[i])
		if i < len(nums)-1 && nums[i+1]-nums[i] == 1 {
			str += "->"
			for i < len(nums)-1 && nums[i+1]-nums[i] == 1 {
				i++
			}
			str += fmt.Sprintf("%d", nums[i])
		}
		res = append(res, str)
	}
	return res
}