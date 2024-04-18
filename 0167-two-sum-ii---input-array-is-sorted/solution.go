func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers) - 1

    for left < len(numbers) - 1 {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left + 1, right + 1}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }

    return []int{0, 0}
}
