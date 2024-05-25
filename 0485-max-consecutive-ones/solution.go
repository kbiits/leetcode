func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    maxCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 1 {
            if count > maxCount {
                maxCount = count
            }
            count = 0
        } else {
            count++
        }
    }

    // for last index
    if count > maxCount {
        maxCount = count
    }

    return maxCount
}
