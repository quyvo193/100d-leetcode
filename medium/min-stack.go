package main

import "fmt"

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if pop == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(-3)
	obj.Push(0)
	obj.Push(-4)
	obj.Push(2)
	fmt.Println("stack: ", obj.stack)

	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()

	min := obj.GetMin()

	fmt.Printf("min: %v\nstack: %v\n", min, obj.stack)
}

// 1 -3 0 -4
// 1 -3 -4
