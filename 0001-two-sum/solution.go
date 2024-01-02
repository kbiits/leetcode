func twoSum(nums []int, target int) []int {
    var mapSum = map[int]int{}
    for i, num := range nums {
        complement := target - num
        if v, ex := mapSum[complement]; ex {
            return []int{i, v}
        } else {
            mapSum[num] = i
        }
    }

    return []int{-1, -1}
}
