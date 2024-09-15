func findMin(nums []int) int {
    left := 0
    right := len(nums) - 1
    for left < right {
        mid := left + (right - left) / 2
        if mid > 0 && nums[mid - 1] > nums[mid] {
            return nums[mid]
        }
        // if mid greater than right, so the min element must be in right part of the array 
        if nums[mid] > nums[right] {
            // search right
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return nums[left]
}

