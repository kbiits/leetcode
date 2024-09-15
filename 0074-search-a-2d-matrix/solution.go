func searchMatrix(matrix [][]int, target int) bool {
    rowX := 0
    rowY := len(matrix) - 1
    for rowX <= rowY {
        midRow := rowX + (rowY - rowX) / 2
        if target < matrix[midRow][0] {
            rowY = midRow - 1
        } else if target > matrix[midRow][len(matrix[midRow]) - 1] {
            rowX = midRow + 1
        } else {
            return findInRow(matrix[midRow], target)
        }
    }

    return false
}

func findInRow(row []int, target int) (bool) {
    left := 0
    right := len(row) - 1
    for left <= right {
        mid := left + (right - left) / 2
        if row[mid] == target {
            return true
        }

        if target < row[mid] {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return false
}
