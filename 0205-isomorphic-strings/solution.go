func isIsomorphic(s string, t string) bool {
    dp1 := map[rune]rune{}
    dp2 := map[rune]rune{}

    for i, c := range s {
        pair := rune(t[i])

        morph, ok := dp1[c]
        if ok && morph != pair {
            return false
        } else if v, ok := dp2[pair]; ok && v != c {
            return false
        } else {
            dp1[c] = pair
            dp2[pair] = c
        }
    }

    return true
}
