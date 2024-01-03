func findWordsContaining(words []string, x byte) []int {
    res := make([]int, 0)
    for i, word := range words {
        found := false
        for _, c := range word {
            if c == rune(x) {
                found = true
                break
            }
        }

        if found {
            res = append(res, i)
        }
    }

    return res
}
