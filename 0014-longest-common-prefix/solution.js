/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    let result = "";
    for (let i = 0; i < strs[0].length; i++) {
        let allSame = true;
        for (let j = 1; j < strs.length; j++) {
            if (i >= strs[j].length || strs[j][i] !== strs[0][i]) {
                allSame = false;
                break;
            }
        }

        if (allSame) {
            result += strs[0][i];
        } else {
            break;
        }
    }

    return result;
};
