func longestConsecutive(nums []int) int {
    // [100,4,200,1,3,2]
    // set :
    // 1 => 1
    // 2 => 1
    // 3 => 1
    // 4 => 1
    // 100 => 1
    // 200 => 1

    set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
        // if found left neighbor, continue.
        // we need to find the start index (element without left neighbor)
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num + 1
		for set[temp] {
			sequence++
			temp++
		}

		res = max(res, sequence)
        if sequence > len(nums) / 2{
            break
        }
	}
	return res
}
