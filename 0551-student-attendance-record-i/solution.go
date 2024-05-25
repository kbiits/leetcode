func checkRecord(s string) bool {
    countConsecutiveL := 0
    countA := 0

    for _, v := range s {
        if v == 'L' {
            countConsecutiveL++
        } else if v == 'A' {
            countA++
            countConsecutiveL = 0
        } else {
            countConsecutiveL = 0
        }

        if countConsecutiveL >= 3 || countA >= 2 {
            return false
        }
    }

    return countA < 2 && countConsecutiveL < 3
}
