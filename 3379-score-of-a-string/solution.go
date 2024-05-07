func scoreOfString(s string) int {
    sum := 0
    for i := 0; i < len(s) - 1; i++ {
        sum += int(absDiff(s[i], s[i + 1]))
    }

    return sum
}

func absDiff(a, b uint8) uint8 {
    if a < b {
        return b - a
    }

    return a - b
}
