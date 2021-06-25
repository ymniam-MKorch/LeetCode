func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += plus
		plus = 0
		if digits[i] >= 10 {
			digits[i] = 0
			plus = 1
		} else {
			break
		}
	}
	if plus == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}