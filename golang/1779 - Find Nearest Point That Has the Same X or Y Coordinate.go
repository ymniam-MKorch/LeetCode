func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func nearestValidPoint(x int, y int, points [][]int) int {
	ans := -1
	distance := 10000
	for i, v := range points {
		if v[0] == x || v[1] == y {
			if abs(v[0]-x)+abs(v[1]-y) < distance {
				distance = abs(v[0]-x) + abs(v[1]-y)
				ans = i
			}
		}
	}
	return ans
}