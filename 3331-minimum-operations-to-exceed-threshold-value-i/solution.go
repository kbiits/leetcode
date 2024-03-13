func minOperations(nums []int, k int) int {
    mapNum := map[int]int{}
    for _, v := range nums {
        mapNum[v]++
    }

    counter := 0
    for key, v := range mapNum {
        if key < k {
            counter += v
        }
    }

    return counter
}
