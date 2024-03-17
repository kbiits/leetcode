func isAcronym(words []string, s string) bool {
    if len(s) != len(words) {
        return false
    }
    
    result := true

    for i, c := range s {
        if c != rune(words[i][0]) {
            result = false
            break
        }
    }

    return result
}
