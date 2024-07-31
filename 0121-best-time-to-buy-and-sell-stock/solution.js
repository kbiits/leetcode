/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let min = 1_00_001;
    let profit = 0;
    
    for (let i = 0; i < prices.length; i++) {
        let price = prices[i];
        if (price < min) {
            min = price;
        }

        let profitTemp = price - min;
        if (profitTemp > profit) {
            profit = profitTemp;
        }
    }

    return profit;
};
