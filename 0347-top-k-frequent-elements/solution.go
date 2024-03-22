func topKFrequent(nums []int, k int) []int {
    // bucket sort by values algorithm


    result := make([]int, 0, k)
    
    theMap := map[int]int{}
    for _, v := range nums {
        theMap[v]++
    }

    bucket := make([][]int, len(nums) + 1)
    for k, v := range theMap {
        bucket[v] = append(bucket[v], k)
    }

    resultCount := 0
    for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
        keys := bucket[i]
        if len(keys) + resultCount > k {
            numNeeded := k - resultCount
            result = append(result, keys[:numNeeded]...)
            resultCount += numNeeded
        } else {
            result = append(result, keys[:]...)
            resultCount += len(keys)
        }
    }

    return result
}
