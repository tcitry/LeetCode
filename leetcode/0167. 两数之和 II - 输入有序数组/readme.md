https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/


## 第一次的独立解法

https://leetcode.cn/submissions/detail/316941840/

还是两数相加的思路，利用哈希表，具有 空间和时间都是 O(n) 的复杂度。

```go
func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)
    for i, n := range numbers {
        diff := target - n
        if v, in := m[diff]; in {
            return []int{v+1,i+1}
        }
        m[n] = i
    }
    return []int{-1, -1}
}
```

## 还有暴力法

两层 for 循环的 时间复杂度 O(n2) 的解法。


## 二分查找

固定第一个数，利用二分法查找第二个数。利用有序属性，决定二分的方向。

- 时间复杂度 O(nlogn)
- 空间复杂度 O(1)


## 双指针

同样利用有序的属性。

1. 两个指针 left、right 分别指向第一个元素和最后一个元素。
2. 如果两个指针之和小于目标值，则 left 右移一位。
3. 如果两个指针之和大于目标值，则 right 左移一位。

这个过程不会跳过答案，时间复杂度 最多 O(n)，空间复杂度 O(1)
