## 题目地址(989. 数组形式的整数加法)

https://leetcode-cn.com/problems/add-to-array-form-of-integer/

## 思路

主要是要构造两个数组相加。

由于不知道 A 的列表长度大还是 K 转换后的列表长度大，所以在 K 转换为数组后，对比后得到一个长数组和短数组

最后针对长数组和短数组做相加运算。

## 代码

Go Code:

```go
func addToArrayForm(A []int, K int) []int {
	var addFlag bool

	var ks []int

	for K/10 > 0 {
		ks = append([]int{K % 10}, ks...)
		K = K / 10
	}
	ks = append([]int{K % 10}, ks...)

	var ll, gl int
	var ls, gs []int
	if len(ks) > len(A) {
		gl = len(ks)
		ll = len(A)
		ls = A
		gs = ks
	} else {
		gl = len(A)
		ll = len(ks)
		ls = ks
		gs = A
	}

	var i, j, item int
	for j = gl - 1; j >= 0; j-- {
		i = j - (gl - ll)
		if i >= 0 {
			item = gs[j] + ls[i]
		} else {
			item = gs[j]
		}
		if addFlag == true {
			item = item + 1
		}
		if item >= 10 {
			gs[j] = item % 10
			addFlag = true
		} else {
			addFlag = false
			gs[j] = item
		}
	}
	if addFlag == true {
		gs = append([]int{1}, gs...)
	}
	return gs
}
```

**复杂度分析**

令 n 为数组长度。

- 时间复杂度：$O(n)$
- 空间复杂度：$O(n)$

执行用时：28 ms, 在所有 Go 提交中击败了 97.84%的用户

内存消耗：6.6 MB, 在所有 Go 提交中击败了 77.35%的用户

## 总结

代码看起来很繁琐，需要看下别人的优秀实现
