func flipAndInvertImage(image [][]int) [][]int {
    for i := 0; i < len(image); i++ {
        for j := 0; j < len(image) / 2; j++ {
            // reverse row and flip the bit
            image[i][j], image[i][len(image) - 1 - j] = image[i][len(image) - 1 - j] ^ 1, image[i][j] ^ 1
        }

        // invert the middle bit
        if len(image) % 2 == 1 {
            image[i][len(image)/2] = image[i][len(image)/2] ^ 1
        }
    }

    return image
}
