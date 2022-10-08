var check = function(s, start, end) {
    
    while(start >= 0 && end < s.length && s[start] === s[end]) {
        start--;
        end++;
    }
    
    return end - start - 1;
}

/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    let maxAnswer = 0;
    let startIdx = 0;
    
    for (let i = 0; i < s.length; i++) {
        
        
        const lenOdd = check(s, i, i);
        const lenEven = check(s, i, i + 1)
        const len = Math.max(lenOdd, lenEven);
        
        
        if (len > maxAnswer) {
            maxAnswer = len;
            startIdx = i - Math.floor((len - 1) / 2);
        }
    }
    
    return s.substring(startIdx, startIdx + maxAnswer);
};
