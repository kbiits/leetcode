func numOfSubarrays(arr []int, k int, threshold int) int {
    count := 0
    windowLength := 0
    sum := 0

    for i := 0; i < len(arr); i++ {
        sum += arr[i]
        windowLength++

        if windowLength == k {
            if sum / k >= threshold {
                count++
            }
            // subtract with first element in window to prepare next loop
            sum -= arr[i - k + 1]
            windowLength--
        }
    }

    return count
}

