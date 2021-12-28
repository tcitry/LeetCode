https://leetcode-cn.com/problems/decode-string/


## 题解

第一次按照“3[a]2[bc]”实现如下，但是发现官方的用例为"3[a2[c]]"，所以不能满足，没有明确理解题意。

```go
func decodeString(s string) string {
	var start, end, times int
	var result, item string
	inFlag := false
	for i := 0; i < len(s); i ++ {
		item = string(s[i])
		if item == "[" {
			result = result[0:len(result)-1]
			start = i + 1
			inFlag = true
			times, _ = strconv.Atoi(string(s[i-1]))
			continue
		}else if item == "]" {
			end = i
			for j := 0; j < times; j ++ {
				result += s[start:end]
			}
			inFlag = false
		}else if inFlag == false {
			result += item
		}
	}
	return result
}

```

后来发现可以靠遍历和递归解决。

## 递归







