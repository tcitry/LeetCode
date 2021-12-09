/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    if(n < 3) {
        return n
    }
    return climbStairs(n-1) + climbStairs(n-2)
};