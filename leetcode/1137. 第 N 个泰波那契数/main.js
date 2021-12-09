/**
 * @param {number} n
 * @return {number}
 */
var tribonacci = function(n) {
    total = 0
    curr = 1
    pre = 1
    prepre = 0
    if (n===0) {
        return 0
    }
    if (n===1) {
        return 1
    }
    if (n===2) {
        return 1
    }
    for (i = 3; i <= n; i++) {
        total = curr + pre +prepre
        prepre = pre
        pre = curr
        curr = total
    }
    return total
};