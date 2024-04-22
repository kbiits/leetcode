
func maximizeSum(nums []int, k int) int {
    maxEl := -1

    for i := 0; i < len(nums); i++ {
        if nums[i] > maxEl {
            maxEl = nums[i]
        }
    }

    // sum of n natural numbers start with a(1) = n * (a(1) + a(n)) / 2
    return k * (maxEl + (maxEl + k - 1)) / 2
}
