func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        res[i] = make([]int, i+1)


        for j := 0; j < i + 1; j++ {
            if j == 0 || j == i {
                res[i][j] = 1
            } else if i > 1 {
                res[i][j] = res[i - 1][j - 1] + res[i - 1][j]
            }
        }
    }

    return res
}
