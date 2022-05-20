/**
 * @param {string} s
 * @return {number}
 */
var myAtoi = function (s) {
    let res = 0
    let sign = 1
    let i = 0
    while (s[i] === ' ') {
        i++
    }
    if (s[i] === '+' || s[i] === '-') {
        sign = s[i] === '+' ? 1 : -1
        i++
    }
    while (s[i] >= '0' && s[i] <= '9') {
        res = res * 10 + s[i].charCodeAt() - '0'.charCodeAt()
        i++
    }
    return sign * res
};

(function () {
    let s = '42'
    console.log(myAtoi(s))
})()
