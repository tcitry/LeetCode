package main

import (
	"fmt"
	"unicode"
)

// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
func main() {
	fmt.Println("result: ", decodeString("3[a2[c]]"))
}

func decodeString(s string) string {
	return ""
}

func decodeString0(s string) string {
	sStack := make([]byte, 0)
	countStack := make([]int, 0)
	stackIndex := make([]int, 0)
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			count = count*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			countStack = append(countStack, count)
			count = 0

			stackIndex = append(stackIndex, len(sStack))
		} else if s[i] == ']' {
			c := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			index := stackIndex[len(stackIndex)-1]
			stackIndex = stackIndex[:len(stackIndex)-1]

			str := string(sStack[index:])
			sStack = sStack[:index]

			for j := 0; j < c; j++ {
				sStack = append(sStack, []byte(str)...)
			}
		} else {
			sStack = append(sStack, s[i])
		}
	}
	return string(sStack)
}

func isDigit(s string) bool {
	return unicode.IsDigit([]rune(s)[0])
}

func stackToString(s string) string {
	return s
}
