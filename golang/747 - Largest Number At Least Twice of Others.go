func dominantIndex(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    var max1,max2 int
    if nums[0] < nums[1] {
        max1 = 1
    } else {
        max2 = 1
    }
    for i := 2; i < len(nums); i++ {
        if nums[i] > nums[max1] {
            max2 = max1
            max1 = i
        } else if nums[i] > nums[max2] {
            max2 = i
        }
    }
    if nums[max1] < nums[max2]*2 {
        return -1
    }
    return max1
}