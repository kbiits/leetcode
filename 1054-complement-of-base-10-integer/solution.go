func bitwiseComplement(n int) int {
    if n == 0 { // edge cases
        return 1
    }

    // 5 => 101
    // 2 => 010
    // ========== +
    // 3 => 111
    temp := n
    // make all bits in temp to 1
    result := 0
    mask := 0
    for temp > 0 {
        complementRightMostBit:= (temp & 1) ^ 1
        result = result | (complementRightMostBit << mask)
        temp = temp >> 1 // drop right-most bit
        mask++
    }

    return result
}
