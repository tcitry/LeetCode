给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。

如果存在则返回 true，不存在返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 迭代+模拟

本解法为最初自己的实现，能用，但在 for 循环的判断中感觉不够好。

```go
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for index, num := range nums {
		var i int
		if index <= k {
			i = 0
		} else {
			if k > index {
				i = k - index
			} else {
				i = index - k
			}
		}
		for j := index; i < j; i++ {
			if (num-nums[i] <= t && num-nums[i] >= 0) || (num-nums[i] >= -t && num-nums[i] <= 0) {
				if j-i <= k {
					return true
				}
			}
		}
	}
	return false
}
```


## 滑动窗口 + 有序集合




