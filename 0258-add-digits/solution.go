func addDigits(num int) int {
    sum := num
    for sum >= 10 {
        temp := sum
        tempSum := 0
        for temp > 0 {
            tempSum += temp % 10
            temp /= 10
        }

        sum = tempSum
    }

    return sum
}
