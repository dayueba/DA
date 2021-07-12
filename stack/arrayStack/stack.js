class Stack{
    constructor() {
        this.arr = []
    }

    isEmpty() {
        return this.arr.length === 0
    }

    getSize() {
        return this.arr.length
    }

    peek() {
        return this.arr[this.arr.length - 1]
    }

    pop() {
        const val = this.arr[this.arr.length - 1]
        this.arr.length -= 1
        return val
    }

    push(val) {
        this.arr.push(val)
    }
}


(function() {
    const stack = new Stack();
    console.log(stack.isEmpty())
    stack.push(1)
    console.log(stack.isEmpty(), stack.peek(), stack.getSize())
    console.log(stack.pop(), stack.isEmpty(), stack.getSize())
    stack.push(1)
    console.log(stack.isEmpty(), stack.peek(), stack.getSize())
})()