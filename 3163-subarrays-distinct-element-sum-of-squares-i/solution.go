func sumCounts(nums []int) int {
    sum := 0
    for i := 0; i < len(nums); i++ {
        theMap := map[int]int{}
        tempVal := 0
        for j := i; j < len(nums); j++ {
            // distinct
            if theMap[nums[j]] == 0 {
                tempVal++
            }

            theMap[nums[j]]++

            sum += tempVal * tempVal
        }
    }

    return sum
}

func findDistinct(nums []int) int {
    theMap := map[int]int{}
    for _, v := range nums {
        theMap[v]++
    }

    return len(theMap)
}

// 1 => 2
// 2 => 1

// 1 => 2
