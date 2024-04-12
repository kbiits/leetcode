import "slices"

func minimumDifference(nums []int, k int) int {
    slices.Sort(nums)
    diffMinimum := 100001

    for i := k - 1; i < len(nums); i++ {
        j := i + 1 - k

        diff := nums[i] - nums[j] 
        if diff < diffMinimum {
            diffMinimum = diff
        }
    }

    return diffMinimum
}
