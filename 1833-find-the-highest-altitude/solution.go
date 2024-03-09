func largestAltitude(gain []int) int {
    alt := 0
    maxAlt := 0
    for _, v := range gain {
        alt = alt + v
        if alt > maxAlt {
            maxAlt = alt
        }
    }

    return maxAlt
}
