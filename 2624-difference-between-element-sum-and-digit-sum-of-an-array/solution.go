func differenceOfSum(nums []int) int {
    var elementSum int
    var digitSum int
    for _, v := range nums {
        digitSum += sumDigit(v)
        elementSum += v
    }

    diff := elementSum - digitSum
    if diff < 0 {
        return -diff
    }

    return diff
}

func sumDigit(n int) int {
    result := 0
    for n > 0 {
        result += n % 10
        n /= 10
    }

    return result
}
