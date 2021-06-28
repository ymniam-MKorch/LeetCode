func findLHS(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	var max int
	for k, v := range m {
		if x, ok := m[k+1]; ok {
			if x+v > max {
				max = x + v
			}
		}
	}
	return max
}