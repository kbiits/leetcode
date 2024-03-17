
func maxProductDifference(nums []int) int {
    firstMax := -1
    secondMax := -2

    firstMin := 10002
    secondMin := 10001

    for _, v := range nums {
        if v > firstMax {
            secondMax = firstMax
            firstMax = v
        } else if v > secondMax {
            secondMax = v
        }

        if v < firstMin {
            secondMin = firstMin
            firstMin = v
        } else if v < secondMin {
            secondMin = v
        }
    }

    return (firstMax * secondMax) - (firstMin * secondMin)
}
