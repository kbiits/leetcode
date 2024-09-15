import (
    // "math"
)
func minEatingSpeed(piles []int, h int) int {
    // binary search the k (speed)

    // define min and maximum possible k
    minK, maxK := 1, getMax(piles)
    for minK < maxK {
        midK := minK + (maxK - minK) / 2
        // feasible check
        if feasible(piles, midK, h) {
            // try lower k
            maxK = midK
        } else {
            // try greater k
            minK = midK + 1
        }
    }

    return minK
}

func feasible(piles []int, k, h int) bool {
    hoursTaken := 0
    for _, pile := range piles {
        hoursTaken += ((pile - 1) / k) + 1
    }

    return hoursTaken <= h
}

func getMax(piles []int) int {
    max := -1
    for _, pile := range piles {
        if pile > max {
            max = pile
        }
    }

    return max
}
