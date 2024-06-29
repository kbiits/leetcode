func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    maxSum := nums[0]
    currSum := nums[0]

    for i := 1; i < len(nums); i++ {
        v := nums[i]
        
        currSum += v
        if v > currSum {
            currSum = v
        }

        if currSum > maxSum {
            maxSum = currSum
        }
    }

    return maxSum
}
