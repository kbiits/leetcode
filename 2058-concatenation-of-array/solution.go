func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums)*2)
    for i, num := range nums {
        ans[i] = num
        ans[i + len(nums)] = num
    }

    return ans
}
