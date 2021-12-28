在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。

如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：

age[y] <= 0.5 * age[x] + 7
age[y] > age[x]
age[y] > 100 && age[x] < 100
否则，x 将会向 y 发送一条好友请求。

注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。

返回在该社交媒体网站上产生的好友请求总数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/friends-of-appropriate-ages
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 分析过程

1. 需要遍历每两两个数，推断类似插入排序的遍历。

```js
var numFriendRequests = function(ages) {
    let total = 0
    let len = ages.length
    ages.sort((a, b) => a - b);
    for (let i = 0; i < len; i++) {
        for (let j = i+1; j < len; j++) {
            x = ages[j]
            y = ages[i]
            if (x * 0.5 + 7 >= y) {
                continue
            }
            if (x === y) {
                total += 1
            }
            total += 1
        }
    }
    return total
};
```

2. 先排序，再双指针。


```js
var numFriendRequests = function(ages) {
    const n = ages.length;
    ages.sort((a, b) => a - b);
    let left = 0, right = 0, ans = 0;
    for (const age of ages) {
        if (age < 15) {
            continue;
        }
        while (ages[left] <= 0.5 * age + 7) {
            ++left;
        }
        while (right + 1 < n && ages[right + 1] <= age) {
            ++right;
        }
        ans += right - left;
    }
    return ans;
};

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/friends-of-appropriate-ages/solution/gua-ling-de-peng-you-by-leetcode-solutio-v7yk/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```

## 总结

这种条件比较多的，不要尝试取反，否则考虑的情况会更复杂。