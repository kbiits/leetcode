import (
    // "sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
    // O(n^2) solution
    // res := make([]int, 0)
    // for i, n1 := range nums {
    //     smallerCount := 0
    //     for j, n2 := range nums {
    //         if i == j {
    //             continue
    //         }

    //         if n2 < n1 {
    //             smallerCount++
    //         }
    //     }

    //     res = append(res, smallerCount)
    // }

    // return res

    // O(n log n) solution
    // numCopy := make([]int, len(nums))
    // copy(numCopy, nums)
    // sort.Ints(nums)
    // mapIth := make(map[int]int)
    // for i, n := range nums {
    //     _, ex := mapIth[n]
    //     if !ex {
    //         mapIth[n] = i
    //     } else {
    //     }
    // }

    // res := make([]int, 0)
    // for _, v := range numCopy {
    //     res = append(res, mapIth[v])
    // }

    // return res

    mapCounter := make([]int, 101)
    for _, v := range nums {
        mapCounter[v]++
    }

    // populate count for smaller elements  
    for i := 1; i < 101; i++ {
        mapCounter[i] += mapCounter[i - 1]
    }

    res := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if (nums[i] == 0) {
            res[i] = 0
        } else {
            res[i] = mapCounter[nums[i] - 1]
        }
    }

    return res
}
