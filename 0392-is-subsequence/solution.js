/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
    let sPtr = 0;
    let tPtr = 0;

    for (; tPtr < t.length; ) {
        if (s[sPtr] === t[tPtr]) {
            sPtr++;
            tPtr++;
        } else {
            tPtr++;
        }
    }

    return sPtr === s.length;
};
