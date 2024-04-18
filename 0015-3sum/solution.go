import "slices"

func threeSum(nums []int) [][]int {
    res := [][]int{}
    if len(nums) < 3 {
        return res
    }

    slices.Sort(nums)
    targetSum := 0

    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
			continue
		}

        // 2 sum solve
        left, right := i + 1, len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right] + nums[i]

            if sum == targetSum {
                res = append(res, []int{nums[left], nums[right], nums[i]})
                left, right = left + 1, right - 1

                for left < right && nums[left] == nums[left-1] {
					left++
				}
                for left < right && nums[right] == nums[right+1] {
					right--
				}
            } else if sum > targetSum {
                right--
            } else {
                left++
            }
        }

    }

    return res
}
