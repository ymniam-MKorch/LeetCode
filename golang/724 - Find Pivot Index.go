func pivotIndex(nums []int) int {
    var totalSum int
    for _,v := range nums {
        totalSum += v
    }
    var leftSum int
    for i := 0; i < len(nums); i++ {
        totalSum -= nums[i]
        if i > 0 {
            leftSum += nums[i-1]
        }
        if totalSum == leftSum {
            return i
        }
    }
    return -1
}