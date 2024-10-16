func maximumTripletValue(nums []int) int64 {
    max := int64(0)
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if val := value(nums, i, j, k); val > 0 && val > max {
                    max = int64(val)
                }
            }
        }
    }

    return max
}

func value(nums []int, i, j, k int) int64 {
    return int64((nums[i] - nums[j]) * nums[k])
}
