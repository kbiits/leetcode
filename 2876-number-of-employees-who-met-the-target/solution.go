func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    count := 0
    for _, v := range hours {
        if v >= target {
            count++
        }
    }

    return count
}
