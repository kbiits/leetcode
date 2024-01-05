func maximumWealth(accounts [][]int) int {
    maxWealth := -1
    for _, customer := range accounts {
        customerWealth := 0
        for _, m := range customer {
            customerWealth += m
        }

        if customerWealth > maxWealth {
            maxWealth = customerWealth
        }
    }

    return maxWealth
}
