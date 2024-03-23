func isSubsequence(s string, t string) bool {
    pointerMatch := 0

    if len(s) == 0 {
        return true
    }

    arr := make([]rune, len(s))
    for i, v := range s {
        arr[i] = v
    }

    for i := 0; i < len(t) && pointerMatch < len(s); i++ {
        v := rune(t[i])
        if arr[pointerMatch] == v {
            pointerMatch++
        }
    }

    return pointerMatch == len(s)
}
