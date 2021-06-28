func removeDuplicates(s string) string {
	stack := []rune{}
	for _, c := range s {
		if len(stack) == 0 || stack[len(stack)-1] != c {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}