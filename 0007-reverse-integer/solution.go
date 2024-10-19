import (
    "math"
)

func reverse(x int) int {
    result := 0
    for ; x != 0; {
        lastDigit := x % 10
        result = result * 10 + lastDigit

        if result > int(math.Pow(float64(2), float64(31)) - 1) {
            return 0
        } else if result < int(math.Pow(float64(-2), float64(31))) {
            return 0
        }

        x /= 10
    }

    return result
}
