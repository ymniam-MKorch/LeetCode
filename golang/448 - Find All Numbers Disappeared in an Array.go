func findDisappearedNumbers(nums []int) []int {
	var res []int
	for _, v := range nums {
		if nums[abs(v)-1] > 0 {
			nums[abs(v)-1] = -nums[abs(v)-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}