package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT)
	fmt.Println("start...")
	go func() {
		switch sig := <-sigs; sig.String() {
		case syscall.SIGINT.String():
			done <- true
		default:
			fmt.Println("unknown signal")
		}
	}()

	for {
		<-done
		fmt.Println("done")
		break
	}
}
