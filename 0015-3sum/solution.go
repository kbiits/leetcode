import "slices"

type pair struct {
    a int
    b int
    c int
}

func threeSum(nums []int) [][]int {
    slices.Sort(nums)

    res := [][]int{}
    resSet := map[pair]int{}

    targetSum := 0

    for i := 0; i < len(nums); i++ {

        ith := nums[i]

        // 2 sum solve
        left := i + 1
        right := len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right] + ith
            if sum == targetSum {
                potentialRes := []int{nums[left], nums[right], ith}
                slices.Sort(potentialRes)

                // check the set
                pair := pair{potentialRes[0], potentialRes[1], potentialRes[2]}
                if _, ok := resSet[pair]; !ok {
                    resSet[pair]++
                    res = append(res, potentialRes)
                }
                
                left++
                right--

            } else if sum > targetSum {
                right--
            } else {
                left++
            }

        }

    }

    return res
}
