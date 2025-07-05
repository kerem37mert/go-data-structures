package main

import "fmt"

func main() {
	stack := newStack()
	stack.push(5)
	stack.push(4)

	fmt.Println(stack.peek().value) // 4
	fmt.Println(stack.pop().value)  // 4
	fmt.Println(stack.pop().value)  // 5
	fmt.Println(stack.pop())        // <nil>
}
