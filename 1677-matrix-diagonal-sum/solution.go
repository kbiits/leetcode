func diagonalSum(mat [][]int) int {
    n := len(mat)
    center := n / 2
    isOdd := n & 1 == 1

    sum := 0
    for i := 0; i < n; i++ {
        row := mat[i]
        primary := row[i]
        secondary := row[n - 1 - i]

        
        sum += primary + secondary
    }

    if isOdd {
        sum -= mat[center][center]
    }

    return sum
}

// [7,3,1,9] => 7 + 9
// [3,4,6,9] => 4 + 6
// [6,9,6,6] => 9 + 6
// [9,5,8,5] => 9 + 5
