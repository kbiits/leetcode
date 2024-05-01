func pivotIndex(nums []int) int {
    // [1,2,3]
    // [1,3,6]
    // [6,5,3]

    leftPrefixSum := make([]int, len(nums))
    leftPrefixSum[0] = nums[0]
    rightPrefixSum := make([]int, len(nums))
    rightPrefixSum[len(nums) - 1] = nums[len(nums) - 1]
    for i := 1; i < len(nums); i++ {
        leftPrefixSum[i] = nums[i] + leftPrefixSum[i - 1]
        
        endIdx := len(nums) - i - 1 
        rightPrefixSum[endIdx] = nums[endIdx] + rightPrefixSum[endIdx + 1]
    }

    for i := 0; i < len(nums); i++ {
        if leftPrefixSum[i] - nums[i] == rightPrefixSum[i] - nums[i] {
            return i
        }
    }

    return -1
}
