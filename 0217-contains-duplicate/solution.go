func containsDuplicate(nums []int) bool {
    var theMap = map[int]int{}
    for _, v := range nums {
        if theMap[v] > 0 {
            return true
        }
        theMap[v]++
    }

    return false
}
