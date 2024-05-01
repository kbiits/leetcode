func searchInsert(nums []int, target int) int {
    return binarySearch(nums, target, 0, len(nums) - 1)
}

func binarySearch(nums []int, target, start, end int) (int) {
    if start > end {
        return start
    }
    
    mid := start + ((end - start) / 2)
    if nums[mid] == target {
        return mid
    }

    if target < nums[mid] {
        return binarySearch(nums, target, start, mid - 1)
    } else {
        return binarySearch(nums, target, mid + 1, end)
    }
}
