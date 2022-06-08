/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function (nums, k) {
  let ret = 0;
  for (let i = 0; i < nums.length; i++) {
    let total = nums[i];
    if (nums[i] === k) {
      ret += 1;
    }
    for (let j = i + 1; j < nums.length; j++) {
      total += nums[j];
      if (total === k) {
        ret += 1;
      }
    }
  }
  return ret;
};
