/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    let result = 0;
    for (let i = s.length - 1; i >= 0; i--) {
        let char = s[i];
        if (i > 0 && mapRoman[char] > mapRoman[s[i - 1]]) {
            result += mapRoman[char] - mapRoman[s[i - 1]];
            i--;
        } else {
            result += mapRoman[char];
        }
    }

    return result;
};

var mapRoman = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}
