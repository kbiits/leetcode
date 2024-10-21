func removeDuplicates(nums []int) int {
    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[i] != nums[j] {
            i++
            copy := nums[j]
            nums[j] = nums[i]
            nums[i] = copy
        }
    }

    return i + 1
}

