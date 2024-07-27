/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let stack = [];
    for (let i = 0; i < s.length; i++) {
        if (s[i] === ')') {
            const last = stack.pop();
            if (last !== '(') {
                return false;
            }
        } else if (s[i] === '}') {
            const last = stack.pop();
            if (last !== '{') {
                return false;
            }
        } else if (s[i] === ']') {
            const last = stack.pop();
            if (last !== '[') {
                return false;
            }
        } else {
            // opening bracket
            stack.push(s[i]);
        }
    }

    return stack.length === 0;
};
