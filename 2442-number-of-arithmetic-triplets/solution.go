func arithmeticTriplets(nums []int, diff int) int {
    counter := 0
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[j] - nums[i] == diff && nums[k] - nums[j] == diff {
                    counter++
                }
            }
        }
    }

    return counter
}
