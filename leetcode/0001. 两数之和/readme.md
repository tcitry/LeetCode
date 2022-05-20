https://leetcode.cn/problems/two-sum/

## 解法一：两层 for 循环暴力排查

第一道题，没啥好说的。

使用 for 循环嵌套的方式，应该叫做暴力法吧。

Python 代码虽然简洁，但是使用了 in 关键字，golang 并不支持，跑完却发现其实 golang 的更快，Python 不仅慢更占内存。

```python
class Solution:
    def twoSum(self, nums, target):
        num_index = dict()
        for i in range(len(nums)):
            if target - nums[i] in num_index:
                return [num_index[target-nums[i]], i]
            num_index[nums[i]] = i
```

```go
func twoSum(nums []int, target int) []int {
	for i, m := range nums {
		for j, n := range nums[i+1:] {
			if m + n == target {
				return []int{i, j+i+1}
			}
		}
	}
	return nil
}
```

## 解法二：map

避免两层嵌套 for 循环，使用 map 存储已经读取的值，在后面通过读取 map 来判断是否在 map 中。

```js
var twoSum = function (nums, target) {
    let map = {};
    for (let i = 0; i < nums.length; i++) {
        let diff = target - nums[i];
        if (map[diff] !== undefined) {
            return [map[diff], i];
        }
        map[nums[i]] = i;
    }
};
```
