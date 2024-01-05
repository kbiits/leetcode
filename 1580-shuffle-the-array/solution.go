func shuffle(nums []int, n int) []int {
    // [1 2 3 4 1 2 3 4]
    //  0 1 2 3 4 5 6 7
    // n = 4
    // (0), (0 + 4 or we call it n), (1), (1 + 4)

    res := make([]int, 0)

    for i := 0; i < n; i++ {
        res = append(res, nums[i], nums[i + n])
    }

    return res
}
