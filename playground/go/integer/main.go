package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	a := 10
	b := strconv.Itoa(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
}