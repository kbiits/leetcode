func arithmeticTriplets(nums []int, diff int) int {
    theMap := map[int]int{}
    counter := 0
    for _, potentialK := range nums {
        jFound := false
        iFound := false

        potentialJ := potentialK - diff
        if potentialJ > 0 { // means k > j
            // check in our table (hashmap)
            if theMap[potentialJ] > 0 {
                jFound = true
            }
        }

        if potentialI := potentialJ - diff; jFound {
            if theMap[potentialI] > 0 { // means j > i
                iFound = true
            }
        }

        if jFound && iFound {
            counter++
        }

        theMap[potentialK]++
    }

    return counter
}
