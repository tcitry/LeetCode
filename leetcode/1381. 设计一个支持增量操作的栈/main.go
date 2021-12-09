package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	fmt.Println(s)
	fmt.Println(len(s), cap(s), s)
	obj := Constructor(10)
	fmt.Println("stack is ", obj.Stack)
	obj.Push(666)
	obj.Push(666)
	obj.Push(666)
	obj.Push(666)
	obj.Push(666)
	fmt.Println("push 666, stack is ", obj.Stack)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	fmt.Println("pop: stack is: ", obj.Stack)
	obj.Increment(3, 666)
	fmt.Println("stack is :", obj.Stack)
}

type CustomStack struct {
	Stack []int
	Cap   int
	Size  int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		Stack: make([]int, maxSize),
		Cap:   maxSize,
		Size:  0,
	}
}

func (this *CustomStack) Push(x int) {
	if this.Size == 0 {
		this.Stack[0] = x
		this.Size += 1
	} else if this.Size < this.Cap {
		this.Stack[this.Size] = x
		this.Size += 1
	}
}

func (this *CustomStack) Pop() int {
	if this.Size == 0 {
		return -1
	}
	ret := this.Stack[this.Size-1]
	this.Size -= 1
	this.Stack[this.Size] = 0
	return ret
}

func (this *CustomStack) Increment(k int, val int) {
	if k > this.Cap {
		k = this.Cap
	}
	for i := this.Size; i < k; i++ {
		this.Stack = append(this.Stack, 0)
	}
	for i := 0; i < k; i++ {
		this.Stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
