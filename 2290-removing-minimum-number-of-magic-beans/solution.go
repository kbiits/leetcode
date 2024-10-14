import (
    "slices"
)

func minimumRemoval(beans []int) int64 {
    sum := 0
    for _, b := range beans {
        sum += b
    }

    ans := sum
    slices.Sort(beans)

    for i := 0; i < len(beans); i++ {
        ans = min(ans, sum - ((len(beans) - i) * beans[i]))
    }

    return int64(ans)
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
