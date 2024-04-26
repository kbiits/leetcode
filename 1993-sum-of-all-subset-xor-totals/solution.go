func subsetXORSum(nums []int) int {
    sum := 0
    res := [][]int{}
    generateSubsetsXorSum(nums, &res, 0, 0, &sum)
    return sum
}

func generateSubsetsXorSum(nums []int, res *[][]int, sumXor int, index int, sum *int) {
    *sum += sumXor

	for i := index; i < len(nums); i++ {
		// recurse
		generateSubsetsXorSum(nums, res, sumXor ^ nums[i], i+1, sum)
	}
}

