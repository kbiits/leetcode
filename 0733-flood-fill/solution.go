func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    visited := make([][]bool, len(image))
    for i := range image {
        visited[i] = make([]bool, len(image[i]))
        for j := range image[i] {
            visited[i][j] = false
        }
    }

    floodFillRecursive(&image, image[sr][sc], sr, sc, color, &visited)
    image[sr][sc] = color

    return image
}

var movements = [][]int{
    // go top
    {-1, 0},
    // go left
    {0, -1},
    // go bottom
    {1, 0},
    // go right
    {0, 1},
}

func floodFillRecursive(image *[][]int, starting, sr int, sc int, color int, visited *[][]bool) {

    (*visited)[sr][sc] = true
    if (*image)[sr][sc] == starting {
        (*image)[sr][sc] = color
    } else {
        return
    }

    for _, move := range movements {
        newSr, newSc := sr + move[0], sc + move[1]

        // check bound idx
        if newSr >= 0 && newSr < len(*image) && 
            newSc >= 0 && newSc < len((*image)[newSr]) &&
            !(*visited)[newSr][newSc] {
            floodFillRecursive(image, starting, newSr, newSc, color, visited)
        }
    }
}
