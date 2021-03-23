### 解题思路
先生成出两个数组，字符串组成的数组，和 c 在字符串中存在的下标组成的数据，通过对比两个数据，进而计算最短距离。

### 性能分析

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了5.66%的用户

### 代码

```golang
func shortestToChar(S string, C byte) []int {
	var cs, ss, result []int
	var min int

	c := int32(C)
	for index, v := range S {
		if v == c {
			cs = append(cs, index)
		}
		ss = append(ss, index)
	}
	if len(cs) == 0 {
		return []int{}
	}
	for _, i := range ss {
		min = len(ss)
		for _, j := range cs {
			if i < j {
				if j-i < min {
					min = j - i
				}
			} else if i == j {
				min = 0
				break
			} else {
				if i-j < min {
					min = i - j
				}
			}
		}
		result = append(result, min)
	}
	return result
}
```

