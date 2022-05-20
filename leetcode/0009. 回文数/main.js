/**
 * @param {number} x
 * @return {boolean}
 */

// todo
var isPalindrome = function(x) {
    if (x < 0) return false
    let str = x.toString()
    let len = str.length
    for (let i = 0; i < len / 2; i++) {
        if (str[i] !== str[len - 1 - i]) return false
    }
    return true
};

(function () {
    let x = 121
    console.log(isPalindrome(x))
})()