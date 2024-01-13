func leftRightDifference(nums []int) []int {
    ans := make([]int, len(nums))
    leftSum := 0
    rightSum := 0

    for _, v := range nums {
        rightSum += v
    }
    for i, v := range nums {
        rightSum -= v
        sum := leftSum - rightSum
        if sum < 0 {
            sum = -1 * sum
        }
        ans[i] = sum
        leftSum += v
    }

    return ans
}
