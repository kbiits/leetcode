// solution with dp
func countBits(n int) []int {

    ans := make([]int, n + 1)
    // count 1 in binary representation
    for i := 0; i <= n; i++ {
        // 0 =>  00 = 0
        // 1 =>  01 = 1
        // 2 =>  10 = 1
        // 3 =>  11 = 2
        // 4 => 100 = 1
        // 5 => 101 = 2
        // 6 => 110 = 2
        // 7 => 111 = 3

        // 111 & 001 = 1
        ans[i] = ans[i >> 1] + (i & 1)
    }

    return ans
}


// solution with naive bitwise
// func countBits(n int) []int {

//     ans := make([]int, n + 1)
//     // count 1 in binary representation
//     for i := 0; i <= n; i++ {
//         // 0 =>  00 = 0
//         // 1 =>  01 = 1
//         // 2 =>  10 = 1
//         // 3 =>  11 = 2
//         // 4 => 100 = 1
//         // 5 => 101 = 2
//         // 6 => 110 = 2
//         // 7 => 111 = 3

//         // 111 & 001 = 1
//         ans[i] = countBitsOne(i)
//     }

//     return ans
// }

// func countBitsOne(n int) int {
//     bitsOne := 0
//     for n != 0 {
//         if n & 1 == 1 {
//             bitsOne++
//         }
//         n >>= 1
//     }

//     return bitsOne
// }

