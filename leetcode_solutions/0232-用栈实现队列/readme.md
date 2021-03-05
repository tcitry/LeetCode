### 解题思路
生成两个栈 A 和 B，一个A用于存结果，另一个B用于 push 时进行交换，交完换存入A

### 代码

```golang
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

```

