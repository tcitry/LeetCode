假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

https://leetcode-cn.com/problems/climbing-stairs/


## 递归

爬到第 x 级台阶的方案数是爬到第 (x-1) 级台阶的方案数和爬到第 (x-2) 级台阶的方案数的和。所以本质上也是斐波那契数列，不同的是需要处理好边界问题。

```js
var climbStairs = function(n) {
    if(n < 3) {
        return n
    }
    return climbStairs(n-1) + climbStairs(n-2)
};
```


## 迭代

递归存在大量重复计算，所以最好还是自底向上迭代。

```go
func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	total, pre, curr := 0, 0, 1
	for i := 0; i < n; i++ {
		total = pre + curr
		pre = curr
		curr = total
	}
	return total
}
```
