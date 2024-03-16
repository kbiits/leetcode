func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }

    return ""
}

func isPalindrome(text string) bool {
    // chars := []rune(text)
    for i := 0; i < len(text) / 2; i ++ {
        ptrLeft := text[i]
        ptrRight := text[len(text) - 1 - i]
        if ptrLeft != ptrRight {
            return false
        }
    }

    return true
}
