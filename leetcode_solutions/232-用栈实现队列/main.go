package main

func main() {

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
func (this *MyQueue) Push(x int)  {
	this.stackA = stackPush(this.stackA, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

}


/** Get the front element. */
func (this *MyQueue) Peek() int {

}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

}

func stackPush(stack []int, value int) []int {
	return append(stack, value)
}

func stackPop(stack []int) int {
	l := len(stack)
	return stack[l-1:l][0]
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
