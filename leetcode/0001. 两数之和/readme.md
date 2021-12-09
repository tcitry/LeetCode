
## 解法

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

