func numSplits(s string) int {
    counter := 0
    leftDp := map[rune]int{}
    rightDp := map[rune]int{}
    for _, c := range s {
        rightDp[c]++
    }

    for i, c := range s {
        if i == len(s) - 1 {
            continue
        }

        leftDp[c]++
        rightDp[c]--

        if rightDp[c] <= 0 {
            delete(rightDp, c)
        }

        if len(leftDp) == len(rightDp) {
            counter++
        }
    }

    return counter
}

func findNumInMap(dp map[rune]int) int {
    count := 0
    for _, v := range dp {
        if v > 0 {
            count++
        }
    }

    return count
}
