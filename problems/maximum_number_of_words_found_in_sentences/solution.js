/**
 * @param {string[]} sentences
 * @return {number}
 */
var mostWordsFound = function(sentences) {
    return sentences.reduce((acc, curr) => {
        let currVal = (curr.trim().match(/( )/g)?.length ?? 0) + 1;
        return Math.max(currVal, acc);
    }, -1)
};