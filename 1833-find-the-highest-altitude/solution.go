func largestAltitude(gain []int) int {
    alt := 0
    maxAlt := 0
    n := len(gain)
    for i := 0; i < n; i++ {
        alt = alt + gain[i]
        if alt > maxAlt {
            maxAlt = alt
        }
    }

    return maxAlt
}
