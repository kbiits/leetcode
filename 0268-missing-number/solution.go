// import "slices"
func missingNumber(nums []int) int {
    // slices.Sort(nums)

    n := len(nums)
    sum := n * (n+1) / 2
    sumArr := 0
    for _, v := range nums {
        sumArr += v
    }

    return sum - sumArr
}
