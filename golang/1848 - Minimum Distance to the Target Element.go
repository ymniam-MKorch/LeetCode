func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getMinDistance(nums []int, target int, start int) int {
	ans := 1000
	for i, num := range nums {
		if num == target {
			ans = Min(ans, Abs(i-start))
		}
	}
	return ans
}