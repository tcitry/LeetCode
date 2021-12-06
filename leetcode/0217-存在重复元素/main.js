/**
 * @param {number[]} nums
 * @return {boolean}
 */

var containsDuplicate = function(nums) {
    m = {}
    for (var v of nums) {
        if (v in m) {
            return true
        } else {
            m[v] = true
        }
    }
    return false
};

function main(params) {
    containsDuplicate([1,2,3])
}

main()
