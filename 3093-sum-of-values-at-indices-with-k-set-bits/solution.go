// func sumIndicesWithKSetBits(nums []int, k int) int {
//     sum := 0
//     for i, v := range nums {
//         // count1s := count1sInBinary(i)

//         // Brian Kernighan’s Algorithm for counting 1's bit (set bit)
//         // in binary representation of a number
//         index := i
//         count1s := 0
//         for index > 0 {
//             count1s++
//             index = index & (index - 1)
//         }

//         if count1s == k {
//             sum += v
//         }
//     }

//     return sum
// }

var mapBits = make([]int, 256)
func populateLookupTableBits() {
    mapBits[0] = 0
    for i := 0; i < 256; i++ {
        mapBits[i] = (i & 1) + mapBits[i / 2]
    }
}

func sumIndicesWithKSetBits(nums []int, k int) int {
    populateLookupTableBits()

    sum := 0
    for i, v := range nums {
        n := i
        countOf1s := mapBits[n & 0xff] + 
            mapBits[(n >> 8) & 0xff] + 
            mapBits[(n >> 16) & 0xff] + 
            mapBits[n >> 24]
        if countOf1s == k {
            sum += v
        }
    }

    return sum
}

