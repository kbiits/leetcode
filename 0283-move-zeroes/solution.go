func moveZeroes(nums []int)  {

    // find first zero idx
    zeroIdx := -1
    for i, v := range nums {
        if v == 0 {
            zeroIdx = i
            break
        }
    }

    if zeroIdx == -1 {
        return
    }

    ptr := zeroIdx + 1
    for ; ptr < len(nums); ptr++ {
        if nums[ptr] != 0 {

            temp := nums[ptr]
            nums[ptr] = nums[zeroIdx]
            nums[zeroIdx] = temp

            zeroIdx++
        }
    }
}

func swap(i, j int, arr []int) {
    
}
