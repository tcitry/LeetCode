/**
 * @param {number[]} ages
 * @return {number}
 */

var numFriendRequests = function(ages) {
    let total = 0
    let len = ages.length
    ages.sort((a, b) => a - b);
    for (let i = 0; i < len; i++) {
        for (let j = i+1; j < len; j++) {
            x = ages[j]
            y = ages[i]
            if (x * 0.5 + 7 >= y) {
                continue
            }
            if (x === y) {
                total += 1
            }
            total += 1
        }
    }
    return total
};

console.log(numFriendRequests([73,106,39,6,26,15,30,100,71,35,46,112,6,60,110]))