func smallestRangeI(nums []int, k int) int {
    max,min := 0,10000
    for _,v := range nums {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    d := max - min
    if d > 2*k {
        return d - 2*k
    }
    return 0
}