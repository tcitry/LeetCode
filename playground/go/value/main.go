package main

import "fmt"

func main()  {
	s := make(map[string]int)
	printString(s)
	fmt.Println(s)
}

func printString(s map[string]int) {
	s["hello"] = 1
	// fmt.Println(s)
}