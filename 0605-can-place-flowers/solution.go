func canPlaceFlowers(flowerbed []int, n int) bool {
    // append 0 to start and end 
    helper := append(append([]int{0}, flowerbed...), 0)

    // sliding window with window size 3, find 3 zeroes
    for i := 1; i < len(helper) - 1; i++ {
        if helper[i - 1] == 0 && helper[i] == 0 && helper[i + 1] == 0 {
            n--
            helper[i] = 1
        }
    }

    return n <= 0
}
