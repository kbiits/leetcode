func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    sum := 0
    for i, v := range arr {
        oddLeft := (i / 2) + 1
        oddRight := (n - 1 - i) / 2 + 1
        evenLeft := (i + 1) / 2
        evenRight := (n - i) / 2

        sum += v * ((oddLeft * oddRight) + (evenLeft * evenRight))
    }

    return sum
}

// 1 => 1  => 1 - 0
// 3 => 2  => 3 - 1
// 5 => 3  => 5 - 2
// 7 => 4  => 7 - 3
// 9 => 5  => 9 - 4

// 2 => 5 => (2 - 0) => 2

