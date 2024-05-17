func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }

    ranges := make([]string, 0)

    left := 0
    for right := left + 1; right < len(nums); right++ {
        if nums[right] - nums[right - 1] > 1 {
            // build range
            if left != right - 1 {
                ranges = append(ranges, fmt.Sprintf("%d->%d", nums[left], nums[right - 1]))
            } else {
                ranges = append(ranges, fmt.Sprintf("%d", nums[left]))
            }
            left = right
            continue
        }
    }

    if left < len(nums) - 1 {
        ranges = append(ranges, fmt.Sprintf("%d->%d", nums[left], nums[len(nums) - 1]))
    } else {
        ranges = append(ranges, fmt.Sprintf("%d", nums[left]))
    }

    return ranges
}
