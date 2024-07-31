/**
 * @param {number[]} nums
 * @return {number}
 */
var findClosestNumber = function(nums) {
    let minVal = -1e6;
    let minDistance = 1e6;
    for (let num of nums) {
        let dNum = findDistance(num);
        if (dNum < minDistance || (dNum === minDistance && num > minVal)) {
            minVal = num;
            minDistance = dNum;
        }
    }

    return minVal;
};

var findDistance = (num) => num < 0 ? (-num) - 0 : num - 0;
