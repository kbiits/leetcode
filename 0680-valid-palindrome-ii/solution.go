func validPalindrome(s string) bool {

    left := 0
    right := len(s) - 1
    for left < right {
        leftChar := s[left]
        rightChar := s[right]
        
        if leftChar != rightChar {
            // check palindrom if we remove the left / right char

            // try to remove left and right
            return isPalindromeHelper(s, left + 1, right) ||
                      isPalindromeHelper(s, left, right - 1)
        }

        left++
        right--
    }

    return true
}

func isPalindromeHelper(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }

        l++
        r--
    }

    return true
}
