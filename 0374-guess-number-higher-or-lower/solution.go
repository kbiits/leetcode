/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    result := -1
    left := 1
    right := n
    var pick int
    for result != 0 {
        if left != right {
            pick = (left + right - 1) / 2
        } else {
            pick = left
        }

        result = guess(pick)
        if result == -1 {
            right = pick - 1
        } else if result == 1 {
            left = pick + 1
        } else {
            break
        }
    }

    return pick
}
