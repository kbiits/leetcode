// O(n^2)
// func countKDifference(nums []int, k int) int {
//     counter := 0
//     for i := 0; i < len(nums); i++ {
//         for j := i + 1; j < len(nums); j++ {
//             diff := nums[i] - nums[j]
//             if diff < 0 {
//                 diff = -diff
//             }

//             if diff == k {
//                 counter++
//             }
//         }
//     }

//     return counter
// }

// using counting sort algorithm (hash map like)
//  to optimize the runtime complexity
func countKDifference(nums []int, k int) int {
    counter := 0
    hashMap := make([]int, 101)

    for _, v := range nums {
        hashMap[v]++
    }

    // k = 2 => complement (5 - 2),
    // [0, 1, 1, 1, 1, 1]
    //  0  1  2  3  4  5

    for i := k; i < 101; i++ {
       counter += hashMap[i] * hashMap[i - k]
    }

    return counter
}

