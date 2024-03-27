func removeElement(nums []int, val int) int {
    count := 0
    i := 0
    counterIterator := 0
    for ; i < len(nums) && counterIterator < len(nums); i++ {
        counterIterator++
        curr := nums[i]

        if curr == val {
            // remove nums[i]
            for j := i; j < len(nums) - 1; j++ {
                nums[j] = nums[j+1]
            }

            i--
        } else {
            count++
        }
    }

    // fmt.Printf("count: %d\n", count)

    return count
}
