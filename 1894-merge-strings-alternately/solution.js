/**
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function(word1, word2) {
    let result = "";
    for (let i = 0; i < word1.length; i++) {
        result += word1[i];

        if (i < word2.length) {
            result += word2[i];
        }
    }

    if (word2.length > word1.length) {
        result += word2.slice(word1.length);
    }

    return result;
};
