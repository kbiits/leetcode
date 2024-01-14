func differenceOfSum(nums []int) int {
    elSum := 0
    digitSum := 0
    for _, v := range nums {
        elSum += v
        curr := v
        for curr > 0 {
            digitSum += curr % 10
            curr = curr / 10
        }
    }

    res := elSum - digitSum
    if res < 0 {
        return -res
    }

    return res
}
