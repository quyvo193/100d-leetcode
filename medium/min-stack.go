package main

import "fmt"

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.min = min(this.min, val)
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(-3)
	obj.Push(0)
	obj.Push(-4)
	fmt.Println("top", obj.Top())
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Printf("p3: %v, p4: %v\nstack: %v\n", param_3, param_4, obj.stack)
}
