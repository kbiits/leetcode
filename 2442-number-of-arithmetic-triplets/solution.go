func arithmeticTriplets(nums []int, diff int) int {
    counter := 0
    hashMap := make([]int, 201)

    // consider this array
    // [1, 2, 3] with diff = 1
    // j should be (k - diff) => (3 - 1) = 2
    // i should be (k - 2 * diff) => (3 - 2) = 1

    for _, potentialK := range nums {
        var jFound, iFound bool

        if potentialK - diff >= 0 {
            jFound = hashMap[potentialK - diff] > 0
        }
        if potentialK - 2 * diff >= 0 {
            iFound = hashMap[potentialK - 2 * diff] > 0
        }
        
        if jFound && iFound {
            counter++
        }

        hashMap[potentialK]++
    }

    return counter
}
