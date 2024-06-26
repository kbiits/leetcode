func plusOne(digits []int) []int {

    carry := false
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] == 9 {
            carry = true
            digits[i] = 0
        } else {
            digits[i]++
            return digits
        }
    }

    if carry {
        digits = append([]int{1}, digits...)
    }

    return digits
}
