func maxProfit(prices []int) int {
    // two pointer
    maxProfit := 0
    left, right := 0, 1

    for right < len(prices) {
        // check if profitable to buy at left, and sell at right
        if prices[left] < prices[right] {
            if prices[right] - prices[left] > maxProfit {
                maxProfit = prices[right] - prices[left]
            }
        } else {
            // set left to right, because right has the lowest price
            left = right
        }
        
        right++
    }

    return maxProfit
}
