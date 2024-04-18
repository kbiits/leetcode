func findDisappearedNumbers(nums []int) []int {
    length := len(nums)
    res := make([]int, length)
    for _, v := range nums {
        res[v-1] = v
    }
    
    j := 0
    for i := 0; i < length; i++ {
        if res[i] == 0 { 
            res[j] = i + 1
            j++
        }
    }

    return res[:j]
}
