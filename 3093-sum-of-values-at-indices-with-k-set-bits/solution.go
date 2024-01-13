func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i, v := range nums {
        count1s := count1sInBinary(i)
        if count1s == k {
            sum += v
        }
    }

    return sum
}

func count1sInBinary(n int) int {
    count := 0
    for n > 0 {
        count++
        n = n & (n - 1)
    }

    return count
}
