func createTargetArray(nums []int, index []int) []int {
    target := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        idx := index[i]
        val := nums[i]

        // [0, 1, 2, 3, 4]
        if len(target) == 0 || idx == len(target) {
            target = append(target, val)
        } else {
            target = append(target[:idx + 1], target[idx:]...)
            target[idx] = val
        }
    }

    return target
}
