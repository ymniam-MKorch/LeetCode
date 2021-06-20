func maxSubArray(nums []int) int {
    var sum int
    maxSum := -100000
    for _,v := range nums {
        sum += v
        if sum > maxSum {
            maxSum = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return maxSum
}