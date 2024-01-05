func smallerNumbersThanCurrent(nums []int) []int {
    res := make([]int, 0)
    for i, n1 := range nums {
        smallerCount := 0
        for j, n2 := range nums {
            if i == j {
                continue
            }

            if n2 < n1 {
                smallerCount++
            }
        }

        res = append(res, smallerCount)
    }

    return res
}
