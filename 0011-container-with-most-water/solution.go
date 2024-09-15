func maxArea(height []int) int {
    if len(height) == 0 {
        return 0
    }

    left := 0
    right := len(height) - 1
    maxArea := 0
    for left < right {
        y := getMin(height[left], height[right])
        x := right - left
        area := x * y
        if area > maxArea {
            maxArea = area
        }

        if height[left] > height[right] {
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
