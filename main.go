package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go testChannel(c)

	c <- 1
	c <- 2

	time.Sleep(time.Second * 10)
	fmt.Println("main done")
}

func testChannel(c <-chan int) {
	for {
		select {
		case m := <-c:
			fmt.Println("get msg:", m)
		}
	}
}
