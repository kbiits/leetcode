func arrangeCoins(n int) int {
    i := 1
    row := 0
    for {
        if n - i < 0 {
            break
        }
        n -= i
        row++
        i++
    }

    // 5 - 1
    // 4 - 2
    // 2 - 3
    return row
}
