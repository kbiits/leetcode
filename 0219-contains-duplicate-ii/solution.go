func containsNearbyDuplicate(nums []int, k int) bool {
    set := map[int]bool{}
    left := 0
    for _, v := range nums {
        // check if v in our set
        if set[v] {
            return true
        }

        // add v to our set
        set[v] = true

        // sliding window, check if our window length > k
        // abs(i - j) always <= k if our window length is <= k
        if len(set) > k {
            delete(set, nums[left])
            left++
        }
    }

    return false
}

func absDiff(i, j int) int {
    if i > j {
        return i - j
    }

    return j - i
}
