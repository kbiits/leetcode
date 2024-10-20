func check(nums []int) bool {
    countNotSorted := 0
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] > nums[i] {
            countNotSorted++
        }
    }

    if nums[0] < nums[len(nums) - 1] {
        countNotSorted++
    }

    return countNotSorted <= 1
}
