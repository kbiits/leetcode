func countKDifference(nums []int, k int) int {
    // theMap := map[int]int{}
    // maximum k is 99
    // maximum nums[i] is 100
    // so the maximum pair value is 100 + 99 = 199, that's why
    // the array size is 200 (max index 199) 
    theMap := make([]int, 200)
    counter := 0

    for _, pair1 := range nums {
        // find where |pair1 - pair2| = k
        // for the absolute difference, each pair will have
        // 2 potential values that leads to k
        // 1. pair1 - k = pair2.1
        // 2. pair1 - (-k) = pair2.2

        pair2A := pair1 - k
        pair2B := pair1 + k

        if pair2A > 0 && pair2A <= 100 {
            if v := theMap[pair2A];  v > 0 {
                counter += v
            }
        }
        
        if pair2B > 0 && pair2B <= 100 {
            if v := theMap[pair2B]; v > 0 {
                counter += v
            }
        }

        theMap[pair1]++
    }

    return counter
}

// [9,3,1,1,4,5,4,9,5,10]

// [9,(3),1,1,(*4),5,4,9,5,10]

// [9,3,1,1,(4),(*5),4,9,5,10]

// [9,(3),1,1,4,5,(*4),9,5,10]
// [9,3,1,1,4,(5),(*4),9,5,10]

// [9,3,1,1,(4),5,4,9,(*5),10]
// [9,3,1,1,4,5,(4),9,(*5),10]

// [(9),3,1,1,4,5,4,9,5,(*10)]
// [9,3,1,1,4,5,4,(9),5,(*10)]

// 9 => 1
// 3 => 1
// 1 => 2
// 4 => 2
