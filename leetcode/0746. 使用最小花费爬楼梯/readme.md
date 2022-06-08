数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 递归

根据状态转移方程可以直接写出递归代码。

```go
func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	return min(helper(cost, l-2), helper(cost, l-1))
}

func helper(cost []int, n int) int {
	if n < 2 {
		return cost[n]
	}
	return min(helper(cost, n-2), helper(cost, n-1)) + cost[n]
}
```


## 自顶向下

针对递归进行优化，缓存计算结果。

```js
/**
 * @param {number[]} cost
 * @return {number}
 */
var minCostClimbingStairs = function (cost) {
  if (cost.length === 0) {
    return 0;
  }
  let dp = [];
  let n = cost.length;
  helper(n, cost, dp);
  return Math.min(dp[n - 2], dp[n - 1]);
};

var helper = function (i, cost, dp) {
  if (i < 2) {
    dp[i] = cost[i];
    return;
  }
  helper(i - 1, cost, dp);
  helper(i - 2, cost, dp);
  dp[i] = Math.min(dp[i - 1], dp[i - 2]) + cost[i];
};
```


## 自底向上

```go
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	l := len(cost)
	return min(dp[l-1], dp[l-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```
## 空间复杂度 O(1)

```go
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	dp := make([]int, 2)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i%2] = min(dp[0], dp[1]) + cost[i]
	}

	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
