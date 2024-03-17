func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    maxLocal := make([][]int, n - 2)

    for i := 0; i < n - 2; i++ {
        maxLocal[i] = make([]int, n-2)
        
        for j := 0; j < n - 2; j++ {
            centerRow := i+1
            centerCol := j+1

            var maxValue int
            for i := -1; i < 2; i++ {
                for j := -1; j < 2; j++ {
                    if v := grid[centerRow + i][centerCol + j]; v > maxValue {
                        maxValue = v
                    }
                }
            }

            maxLocal[i][j] = maxValue

        }

    } 

    return maxLocal
}

// func largestLocal(grid [][]int) [][]int {
//     n := len(grid)
//     maxLocal := make([][]int, n - 2)

//     for i := 0; i < n - 2; i++ {
//         maxLocal[i] = make([]int, n-2)
        
//         for j := 0; j < n - 2; j++ {
//             centerGridRow := i+1
//             centerGridCol := j+1

//             maxLocal[i][j] = findMaxValIn3x3(centerGridRow, centerGridCol, grid)
//         }

//     } 

//     return maxLocal
// }

// func findMaxValIn3x3(centerRow, centerCol int, grid [][]int) int {
//     var maxValue int
//     for i := -1; i < 2; i++ {
//         for j := -1; j < 2; j++ {
//             if v := grid[centerRow + i][centerCol + j]; v > maxValue {
//                 maxValue = v
//             }
//         }
//     }

//     return maxValue
// }
