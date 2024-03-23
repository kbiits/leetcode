func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n == 1 {
        return strs[0]
    }

    prefix := []rune(strs[0])
    for i := 1; i < n; i++ {
        str := strs[i]

        for j := 0; j < len(prefix); j++ {
            if j >= len(str) || rune(str[j]) != prefix[j] {
                prefix = prefix[:j]
                break
            }
        }
    }

    return string(prefix)
}
