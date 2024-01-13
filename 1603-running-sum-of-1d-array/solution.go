func runningSum(nums []int) []int {
    result := make([]int, len(nums))
    for i, v := range nums {
        if i == 0 {
            result[i] = v
        } else {
            result[i] = result[i - 1] + v
        }
    }

    return result
}
