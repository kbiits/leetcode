func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    theMap := map[rune]int{}
    for _, v := range s {
        theMap[v]++
    }

    theMap2 := map[rune]int{}
    for _, v := range t {
        theMap2[v]++
    }

    for k, v := range theMap {
        if theMap2[k] != v {
            return false
        }
    }

    return true
}
