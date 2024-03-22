func productExceptSelf(nums []int) []int {
    // prefix multiplication (pre-computed)

    if (len(nums) == 0) {
        return []int{}
    }

    if (len(nums) == 1) {
        return []int{nums[0]}
    }

    // nums         [1,  2,  3,   4]
    // prefixSum    [1,  2,  6,  24]     
    // postfixSum   [4, 12, 24,  24] // start from end of array

    prefix := make([]int, len(nums))
    postfix := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        pairEndIdx := len(nums) - 1 - i

        prefixSumBeforeI := 1
        postfixSumAfterI := 1

        if i > 0 {
            prefixSumBeforeI = prefix[i - 1]
            postfixSumAfterI = postfix[i - 1]
        }

        curr := nums[i]
        pairCurr := nums[pairEndIdx]

        prefix[i] = prefixSumBeforeI * curr
        postfix[i] = postfixSumAfterI * pairCurr
    }

    // fmt.Println(prefix)
    // fmt.Println(postfix)

    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        prefixSum := 1
        postfixSum := 1

        postfixIdx := len(nums) - 1 - i
        
        if i > 0 {
            prefixSum = prefix[i-1]
        }

        if postfixIdx > 0 {
            postfixSum = postfix[postfixIdx - 1]
        }
        
        result[i] = prefixSum * postfixSum
    }
    
    return result
}
