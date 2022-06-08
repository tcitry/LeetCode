/**
 * @param {number[]} nums
 * @return {number[][]}
 */

// 第一次方案，未考虑重复的情况
// var threeSum = function (nums) {
//   let ret = [];
//   nums.sort((a, b) => {
//     return a - b;
//   });
//   for (let i = 0; i < nums.length; i++) {
//     for (let j = i + 1; j < nums.length; j++) {
//       let diff = 0 - nums[i] - nums[j];
//       let index = nums.indexOf(diff);
//       if (index !== -1 && index > j && index < nums.length) {
//         ret.push([nums[i], nums[j], diff]);
//       }
//     }
//   }
//   return ret;
// };

// todo
var threeSum = function (nums) {
    let ret = [];
    nums.sort((a, b) => {
        return a - b;
    });
    for (let i = 0; i < nums.length; i++) {
        if (i > 0 && nums[i] === nums[i - 1]) {
            continue;
        }
        for (let j = i + 1; j < nums.length; j++) {
            if (j > i + 1 && nums[j] === nums[j - 1]) {
                continue;
            }
            let diff = 0 - nums[i] - nums[j];
            let index = nums.indexOf(diff);
            if (index !== -1 && index > j && index < nums.length) {
                ret.push([nums[i], nums[j], diff]);
            }
        }
    }
    return ret;
};