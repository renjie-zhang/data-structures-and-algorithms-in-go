package ArrayStack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := &Stack{
		MaxTop: 10,
		Top:    -1,
	}
	stack.Push(23)
	stack.Push(56)
	stack.Push(89)
	stack.Push(89)
	stack.Push(89)
	stack.Push(89)
	stack.Push(89)
	stack.Push(89)
	stack.Push(89)
	stack.Push(96)
	err := stack.Push(89)
	fmt.Println(err)
	val, _ := stack.Pull()
	fmt.Printf("val=%d", val)
}

/*
=== RUN   TestArrayStack
error : stack is full
val=96--- PASS: TestArrayStack (0.00s)
PASS
ok      Data-Structures-and-Algorithms-Go/ArrayStack    0.292s
*/
