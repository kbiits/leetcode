func isPalindrome(s string) bool {
    left := 0
    right := len(s) - 1
    for left < right {
        leftChar := s[left]
        rightChar := s[right]


        if !(leftChar >= 'a' && leftChar <= 'z') &&
            !(leftChar >= 'A' && leftChar <= 'Z') && 
            !(leftChar >= '0' && leftChar <= '9') {
            left++
            continue
        }

        if !(rightChar >= 'a' && rightChar <= 'z') &&
            !(rightChar >= 'A' && rightChar <= 'Z') && 
            !(rightChar >= '0' && rightChar <= '9') {
            right--
            continue
        }

        if leftChar >= 'A' && leftChar <= 'Z' {
            leftChar += 32
        }

        if rightChar >= 'A' && rightChar <= 'Z' {
            rightChar += 32
        }

        if leftChar != rightChar {
            return false
        }

        left++
        right--
    }

    return true
}
