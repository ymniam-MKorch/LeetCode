func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	var res []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok && m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}