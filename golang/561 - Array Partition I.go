func arrayPairSum(nums []int) int {
	arr := make([]int, 20001)
	for _, v := range nums {
		arr[v+10000]++
	}
	var sum int
	flag := true
	for i := 0; i < len(arr); i++ {
		for arr[i] > 0 {
			if flag {
				sum += (i - 10000)
			}
			flag = !flag
			arr[i]--
		}
	}
	return sum
}