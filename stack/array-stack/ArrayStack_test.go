/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-09-04 21:51:30
 * @LastEditTime: 2019-10-21 22:51:16
 */
package ArrayStack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := NewArrayStack(5)
	stack.Push(23)
	stack.Push(56)
	stack.Push(87)
	stack.Push(84)
	stack.Push(87)
	err := stack.Push(89)
	fmt.Println(err)
	val, _ := stack.Pull()
	fmt.Printf("val=%d\n", val)
}

/*
=== RUN   TestArrayStack
error : stack is full
val=96--- PASS: TestArrayStack (0.00s)
PASS
ok      Data-Structures-and-Algorithms-Go/ArrayStack    0.292s
*/
