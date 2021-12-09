/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    if (nums.length === 1) {
        return nums[0]
    }
    let max = nums[0]
    for(i = 1; i < nums.length; i++) {
        if (nums[i] < nums[i] + nums[i-1]) {
            nums[i] += nums[i-1]
        }
        if (max < nums[i]) {
            max = nums[i]
        }
    }
    return max
};