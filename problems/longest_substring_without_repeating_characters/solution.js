/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    let map = {};
    let startWindow = 0;
    let endWindow = 0;
    let maxLength = 0;
    
    while (endWindow < s.length) {
        // make window doesn't contain duplicates
        while (s[endWindow] in map) {
            delete map[s[startWindow]];
            startWindow++;
        }

        map[s[endWindow]] = 1;
        maxLength = Math.max(maxLength, endWindow - startWindow + 1);
        endWindow++;
    }
    
    return maxLength;
};