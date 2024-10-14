import (
    "slices"
)

func heightChecker(heights []int) int {
    originHeights := make([]int, 0, len(heights))
    for _, h := range heights {
        originHeights = append(originHeights, h)
    }

    slices.Sort(heights)

    count := 0
    for i, v := range heights {
        if originHeights[i] != v {
            count++
        }
    }

    return count
}
