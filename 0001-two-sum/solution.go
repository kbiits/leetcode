func twoSum(nums []int, target int) []int {
    
    theMap := map[int]int{}
    for i, v := range nums {
        complement := target - v
        complementIdx, ok := theMap[complement]
        if ok {
            return []int{complementIdx, i}
        }

        theMap[v] = i
    }

    return []int{}
}
