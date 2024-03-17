func countConsistentStrings(allowed string, words []string) int {
    theMap := make([]int, 26)
    for _, v := range allowed {
        theMap[rune(v) - 'a']++
    }

    counter := 0
    for _, word := range words {
        isConsistent := true
        for _, c := range word {
            if theMap[rune(c) - 'a'] <= 0 {
                isConsistent = false
                break
            }
        }

        if isConsistent {
            counter++
        }
    }

    return counter
}
