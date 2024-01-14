func truncateSentence(s string, k int) string {
    if k == 0 {
        return s
    }

    arrStr := []rune(s)
    i := 0
    for i < len(s) {
        if arrStr[i] == ' ' {
            k--
        }

        if k == 0 {
            break
        }

        i++
    }

    return string(arrStr[:i])
}
