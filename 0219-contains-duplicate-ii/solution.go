func containsNearbyDuplicate(nums []int, k int) bool {
    slidingWindow := map[int]bool{}
    left := 0
    for _, v := range nums {
        if _, ok := slidingWindow[v]; ok {
            return true
        }

        slidingWindow[v] = true

        // sliding window
        if len(slidingWindow) > k {
            delete(slidingWindow, nums[left])
            left++
        }
    }

    return false
}

