package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(4)
	obj.Push(3)
	obj.Push(2)
	obj.Push(1)
	param_3 := obj.Peek()
	param_2 := obj.Pop()
	param_21 := obj.Pop()
	param_22 := obj.Pop()
	param_23 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4, param_21, param_22, param_23)
}

type MyQueue struct {
	stackA []int
	stackB []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackA: []int{},
		stackB: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for la := len(this.stackA); la > 0; la-- {
		v := this.stackA[la-1]
		this.stackA = this.stackA[0 : la-1]
		this.stackB = append(this.stackB, v)
	}

	this.stackA = append(this.stackA, x)
	for lb := len(this.stackB); lb > 0; lb-- {
		v := this.stackB[lb-1]
		this.stackB = this.stackB[0 : lb-1]
		this.stackA = append(this.stackA, v)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ret := this.stackA[len(this.stackA)-1]
	this.stackA = this.stackA[0 : len(this.stackA)-1]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stackA[len(this.stackA)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackA) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
