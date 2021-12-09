斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给你 n ，请计算 F(n) 。

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fibonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 递归

```js
var fib = function(n) {
    if(n < 2) {
        return n
    }
    return fib(n-1) + fib(n-2)
};
```

## 迭代

```ts
function fib(n: number): number {
    if (n < 2) {
        return n
    }
    let total: number = 0
    let pre: number = 0
    let curr: number = 1
    for (let i=1; i < n; i++) {
        total = pre + curr
        pre = curr
        curr = total
    }
    return total
};
```