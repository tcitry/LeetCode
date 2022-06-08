https://leetcode.cn/problems/QTMn0o/

方案一 枚举

```js
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
```

方案二 前缀和+哈希表

```js
var subarraySum = function(nums, k) {
    const mp = new Map();
    mp.set(0, 1);
    let count = 0, pre = 0;
    for (const x of nums) {
        pre += x;
        if (mp.has(pre - k)) {
            count += mp.get(pre - k);
        }
        if (mp.has(pre)) {
            mp.set(pre, mp.get(pre) + 1);
        } else {
            mp.set(pre, 1);
        }
    }
    return count;
};

```