/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function (nums, target) {
    let low = 0;
    let high = nums.length - 1;
    while (low <= high) {
        const mid = Math.floor((high - low) / 2) + low;
        const m = nums[mid];
        if (m === target) {
            return mid;
        } else if (m < target) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    return -1;
};
