func countPairs(nums []int, target int) int {
    counter := 0
    for i, n1 := range nums {
        for j := i + 1; j < len(nums); j++ {
            n2 := nums[j]

            if n1 + n2 < target {
                counter++
            }
        }
    }

    return counter
}
