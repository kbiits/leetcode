func removeDuplicates(nums []int) int {
    pointer := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i - 1] {
            nums[pointer] = nums[i]
            pointer++
        }
    }

    return pointer
}
