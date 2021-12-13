/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    const pattern = "abc";
    while(s.includes("abc")) {
        s = s.replace(/abc/g, "");
    }
    
    return s.length === 0;
};