func intersection(nums1 []int, nums2 []int) []int {
    memory := map[int]int{}
    for _, v := range nums1 {
        memory[v] = 2
    }

    res := []int{}
    for _, v := range nums2 {
        if val := memory[v]; val == 2 {
            res = append(res, v)
            memory[v]--
        }
    }

    return res
}
