### 解题思路
使用二分法，具体怎么样的二分，一般是 (l + r + 1) / 2，也可以 (r - l + 1) / 2。

另外就是注意 l 和 r 的边界值，保证循环能够结束。

### 代码

```golang
func mySqrt(x int) int {
	l := 0
	r := x
	mid := 0
	for ; l < r ; {
		mid = (l + r + 1) / 2
		if mid * mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}
```
