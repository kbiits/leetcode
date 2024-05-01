func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    reversed := 0
    xCopy := x
    for xCopy != 0 {
        reversed *= 10
        reversed += xCopy % 10
        xCopy /= 10
    }

    return reversed == x
}
