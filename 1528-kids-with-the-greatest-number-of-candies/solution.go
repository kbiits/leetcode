func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandiesKidHas := 0
    for _, c := range candies {
        if c > maxCandiesKidHas {
            maxCandiesKidHas = c
        }
    }

    res := make([]bool, 0)
    for _, c := range candies {
        if c + extraCandies >= maxCandiesKidHas {
            res = append(res, true)
        } else {
            res = append(res, false)
        }
    }

    return res
}
