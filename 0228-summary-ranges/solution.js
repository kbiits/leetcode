/**
 * @param {number[]} nums
 * @return {string[]}
 */
var summaryRanges = function(nums) {
    if (nums.length === 0) return [];

    let result = [];
    let leftPtr = 0;
    for (let rightPtr = leftPtr + 1; rightPtr < nums.length; rightPtr++) {
        if (nums[rightPtr] - nums[rightPtr - 1] > 1) { // stop the range
            if (leftPtr !== rightPtr - 1) {
                result.push(`${nums[leftPtr]}->${nums[rightPtr - 1]}`);
            } else {
                result.push(`${nums[leftPtr]}`);
            }
            leftPtr = rightPtr;
            continue;
        }
    }

    if (leftPtr < nums.length - 1) { // left is not the latest element
        result.push(`${nums[leftPtr]}->${nums[nums.length - 1]}`);
    } else {
        result.push(`${nums[leftPtr]}`);
    }

    return result;
};
