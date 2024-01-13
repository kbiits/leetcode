func decompressRLElist(nums []int) []int {
    res := make([]int, 0)
    for i := 0; i < len(nums) / 2; i++ {
        freq := nums[2 * i]
        val := nums[2 * i + 1]
        
        for j := 0; j < freq; j++ {
            res = append(res, val)
        }
    }

    return res
}
