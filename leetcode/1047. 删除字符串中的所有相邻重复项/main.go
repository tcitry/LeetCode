package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

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
