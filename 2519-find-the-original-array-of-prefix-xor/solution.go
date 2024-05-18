func findArray(pref []int) []int {
    arr := make([]int, len(pref))

    temp := pref[0]
    arr[0] = temp
    for i := 1; i < len(pref); i++ {
        res := temp ^ pref[i]
        arr[i] = res
        temp ^= res
    }

    return arr
}

// 101
// 110
// 010
// ==== ^
// 001

