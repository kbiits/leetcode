func twoSum(nums []int, target int) []int {
	var table = map[int]int{}
	var result = []int{0, 0}
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		complement := target - curr
		if v, ex := table[complement]; ex {
			result[0] = v
            result[1] = i
            break
		}

		table[curr] = i
	}

	return result
}
