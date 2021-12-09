泰波那契序列 Tn 定义如下： 

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

提示：

0 <= n <= 37
答案保证是一个 32 位整数，即 answer <= 2^31 - 1。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-th-tribonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



## 递归

递归能实现，但是 LeetCode 判定为超时，所以最好还是迭代。

```go
func tribonacci(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	if n <= 3 {
		return 2
	}
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}
```


## 迭代

```js
var tribonacci = function(n) {
    total = 0
    curr = 1
    pre = 1
    prepre = 0
    if (n===0) {
        return 0
    }
    if (n===1) {
        return 1
    }
    if (n===2) {
        return 1
    }
    for (i = 3; i <= n; i++) {
        total = curr + pre +prepre
        prepre = pre
        pre = curr
        curr = total
    }
    return total
};
```