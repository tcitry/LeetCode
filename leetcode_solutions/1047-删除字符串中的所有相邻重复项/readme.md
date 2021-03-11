### 解题思路

递归，这种方法最简答，但是性能最差。

### 代码

```go
func removeDuplicates(S string) string {
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			S = S[0:i-1] + S[i+1:]
			S = removeDuplicates(S)
			break
		}
	}
	return S
}
```
