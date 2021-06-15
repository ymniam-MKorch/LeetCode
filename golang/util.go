package util

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
