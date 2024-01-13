func restoreString(s string, indices []int) string {
    val := make([]rune, len(s))

    for i := 0; i < len(s); i++ {
        index := indices[i]
        val[index] = rune(s[i])
    }

    return string(val)
}
