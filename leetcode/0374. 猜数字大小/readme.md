https://leetcode.cn/problems/guess-number-higher-or-lower/


自己的二分法

```js
var guessNumber = function (n) {
  let low = 1;
  let high = n;
  while (low < high) {
    let mid = Math.floor((high - low) / 2) + low;
    let ret = guess(mid);
    if (ret === 0) {
      return mid;
    } else if (ret === -1) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }
  return low;
};
```

跟标准答案还是有差距

```js
var guessNumber = function(n) {
    let left = 1, right = n;
    while (left < right) { // 循环直至区间左右端点相同
        const mid = Math.floor(left + (right - left) / 2); 
        if (guess(mid) <= 0) {
            right = mid; // 答案在区间 [left, mid] 中
        } else {
            left = mid + 1; // 答案在区间 [mid+1, right] 中
        }
    }
    // 此时有 left == right，区间缩为一个点，即为答案
    return left;
};
```