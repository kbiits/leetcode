func majorityElement(nums []int) int {
    candidate := nums[0]
    count := 1

    // boyer-moore voting algorithms
    for i := 1; i < len(nums); i++ {
        v := nums[i]
        if count == 0 {
            candidate = v
        }

        if v == candidate {
            count++
        } else {
            count--
        }
    }

    return candidate
}
