func numIdenticalPairs(nums []int) int {
    counter := 0
    mapPair := map[int]int{}
    for _, num := range nums {
        mapPair[num]++
    }

    for _, v := range mapPair {
        counter += (v * (v - 1)) / 2
    }

    return counter
}
