func minPartitions(n string) int {
    maxDigit := -1
    for _, c := range n {
        if int(c - '0') > maxDigit {
            maxDigit = int(c - '0')
        }
    }

    return maxDigit
}
