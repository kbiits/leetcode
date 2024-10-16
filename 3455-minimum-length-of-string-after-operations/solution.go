func minimumLength(s string) int {
    type val struct {
        count int
    }
    dp := make([]*val, 26)

    countRemoved := 0
    for _, c := range s {
        idx := c - 'a'
       v := dp[idx]
       if v == nil {
        v = &val{}
        dp[idx] = v
       }
       
       count := v.count + 1
       if count == 3 {
            countRemoved++
            count -= 2
       }
       v.count = count
       dp[idx] = v
    }

    return len(s) - (countRemoved * 2)
}
