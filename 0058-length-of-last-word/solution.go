func lengthOfLastWord(s string) int {
    countChars := 0
    
    for i := len(s) - 1; i >= 0; i-- {
        chr := rune(s[i])
        if countChars == 0 && chr == ' ' {
            continue
        } else if chr == ' ' {
            break
        } else {
            countChars++
        }
    }

    return countChars
}
