func minimumOperations(nums []int) int {
    count := 0
    for _, num := range nums {
        mod := num % 3
        modCeil := 3 - mod
        if mod < modCeil {
            count += mod
        } else {
            count += modCeil
        }
    }

    return count
}
