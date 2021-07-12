package stack

type Stack interface {
	push(data int)
	pop() int
	peek() int
	getSize() int
	isEmpty() bool
}
