import (
    "sort"
)

func numberGame(nums []int) []int {
    sort.Ints(nums)

    // this idea is to sort the nums, and then reconstruct the result array 
    result := make([]int, len(nums))
    for i := 0; i < len(nums) / 2; i++ {
        // 0 => 0, 1
        // 1 => 2, 3
        // 2 => 4, 5
        // 3 => 6, 7
        // [2, 3, 4, 5]
        startIndex := i * 2
        result[startIndex] = nums[startIndex + 1]
        result[startIndex + 1] = nums[startIndex]
    }

    return result
}
