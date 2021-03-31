### 解题思路

https://leetcode-cn.com/problems/design-a-stack-with-increment-operation/

还是数组问题，实现一个栈

### 解题思路
构造一个 struct，然后操作 push 和 pop，注意 size 的边界。

```go
type CustomStack struct {
	Stack []int
	Cap   int
	Size  int
}
```

涉及到的基础，对 go 的 array 和 slice 有一个区别的认识。

### 代码

```golang

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
```

这是一个别人的解决方案，对比起来占用内容更小，学习一下。。本质上是我 make 的使用姿势不对，对第二个参数 length 设为0 就可以了，不需要再初始化的时候设置每一位默认都是 0。

```go
type CustomStack struct {
    Stack []int
}
 
func Constructor(maxSize int) CustomStack {
    return CustomStack{
        Stack: make([]int, 0, maxSize),
    }
}
 
func (this *CustomStack) Push(x int) {
    if len(this.Stack) < cap(this.Stack) {
        this.Stack = append(this.Stack, x)
    }
}
 
func (this *CustomStack) Pop() int {
    if len(this.Stack) == 0 {
        return -1
    }
    res := this.Stack[len(this.Stack)-1]
    this.Stack = this.Stack[:len(this.Stack)-1]
    return res
}
 
func (this *CustomStack) Increment(k int, val int) {
    if len(this.Stack) < k {
        k = len(this.Stack)
    }
    for i := 0; i < k; i++ {
        this.Stack[i] += val
    }
}
```
