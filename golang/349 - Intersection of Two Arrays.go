func intersection(nums1 []int, nums2 []int) []int {
    map1 := make(map[int]bool,0)
    for _, v := range nums1 {
        map1[v] = true
    }
    resMap := make(map[int]bool,0)
    for _, v := range nums2 {
        if _,ok := map1[v]; ok {
            resMap[v] = true
        }
    }
    res := make([]int, 0, len(resMap))
    for k,_ := range resMap {
        res = append(res, k)
    }
    return res
}

// 删map
func intersection2(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}