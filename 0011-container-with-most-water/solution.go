func maxArea(height []int) int {
    if len(height) == 0 {
        return 0
    }

    left := 0
    right := len(height) - 1
    maxArea := 0
    for left < right {
        leftY := height[left]
        rightY := height[right]
        y := getMin(leftY, rightY)
        x := right - left
        area := y * x
        if area > maxArea {
            maxArea = area
        }

        if leftY >= rightY {
            right--
        } else {
            left++
        }
    }

    return maxArea
}

func getMin(a, b int) int {
    if a <= b {
        return a
    }

    return b
}
