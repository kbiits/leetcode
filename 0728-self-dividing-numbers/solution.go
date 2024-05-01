func selfDividingNumbers(left int, right int) []int {
    res := []int{}
    for i := left; i <= right; i++ {
        flag := true
        iCopy := i

        for iCopy != 0 {
            digit := iCopy % 10
            if digit == 0 {
                flag = false
                break
            }

            if i % digit != 0 {
                flag = false
                break
            }

            iCopy /= 10
        }

        if flag {
            res = append(res, i)
        }
    }

    return res
}
