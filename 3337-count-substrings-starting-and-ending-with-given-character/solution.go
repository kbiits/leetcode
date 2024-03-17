func countSubstrings(s string, c byte) int64 {
    total := 0
    for _, v := range s {
        if v == rune(c) {
            total++
        }
    }

    return int64(total * (total + 1) / 2)
}
