package lc1869

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func checkZeroOnes(s string) bool {
	var maxLen0, maxLen1 int

	for i := 0; i < len(s); {
		var len0, len1 int
		if i < len(s) && s[i] == '0' {
			len0++
			i++
			for i < len(s) && s[i] != '1' {
				len0++
				i++
			}
			maxLen0 = Max(maxLen0, len0)
		}
		if i < len(s) && s[i] == '1' {
			len1++
			i++
			for i < len(s) && s[i] != '0' {
				len1++
				i++
			}
			maxLen1 = Max(maxLen1, len1)
		}
	}

	return maxLen1 > maxLen0
}

func main() {
	a := "111000"
	println(checkZeroOnes(a))
}
