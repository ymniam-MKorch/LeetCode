func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res, gi, si := 0, 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			gi++
			si++
		} else {
			si++
		}
	}
	return res
}