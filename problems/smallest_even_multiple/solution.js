/**
 * @param {number} n
 * @return {number}
 */
var smallestEvenMultiple = function(n) {
    if (n % 2 == 0) {
        return n;
    }
    
    return lcm(n, 2);
};

var lcm = function (a, b) {
    return (a / gcd(a, b)) * b;
}

var gcd = function (a, b) {
    if (a == 0) {
        return b;
    }
    
    return gcd(b % a, a);
}