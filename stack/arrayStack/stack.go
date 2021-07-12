package main

import "fmt"

type ArrayStack struct {
	arr []int
}

func (stack *ArrayStack) push(data int) {
	stack.arr = append(stack.arr, data)
}

func (stack *ArrayStack) pop() int {
	length := len(stack.arr)
	val := stack.arr[length-1]
	stack.arr = stack.arr[:length-1]
	return val
}

func (stack *ArrayStack) peek() int {
	return stack.arr[len(stack.arr)-1]
}

func (stack *ArrayStack) getSize() int {
	return len(stack.arr)
}

func (stack *ArrayStack) isEmpty() bool {
	return len(stack.arr) == 0
}

func main() {
	stack := ArrayStack{}
	fmt.Println(stack.isEmpty())
	stack.push(1)
	fmt.Println(stack.isEmpty(), stack.peek(), stack.getSize())
	fmt.Println(stack.pop(), stack.isEmpty())
}
