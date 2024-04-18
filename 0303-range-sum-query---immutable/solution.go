type NumArray struct {
    nums []int
}

// 1, 2, 3, 4
// 1, 3, 6, 10


func Constructor(nums []int) NumArray {
    prefixSum := make([]int, len(nums))
    if len(nums) > 0 {
        prefixSum[0] = nums[0]
    }

    for i := 1; i < len(nums); i++ {
        prefixSum[i] = nums[i] + prefixSum[i - 1]
    }

    return NumArray{
        nums: prefixSum,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.nums[right]
    }

    return this.nums[right] - this.nums[left - 1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
