import (
    "sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
    x := make([]int, 0)
    for _, point := range points {
        x = append(x, point[0])
    }

    sort.Ints(x)

    currentMax := -1
    for i := 0; i < len(x) - 1; i++ {
        v1 := x[i]
        v2 := x[i + 1]
        diff := v2 - v1
        if diff > currentMax {
            currentMax = diff
        }
    }

    return currentMax
}
