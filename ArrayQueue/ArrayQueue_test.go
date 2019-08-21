package ArrayQueue

import (
	"testing"
)

/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-21 19:10:58
 * @LastEditTime: 2019-08-21 19:21:15
 */
func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(6)
	q.EnQueue(5)
	q.EnQueue(4)
	q.EnQueue(3)
	q.EnQueue(2)
	t.Log(q)
	q.DeQueue()
	q.DeQueue()
	t.Log("取出两个元素之后")
	t.Log(q)
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	t.Log("全部取出之后")
	t.Log()
}

/*
=== RUN   TestArrayQueue
--- PASS: TestArrayQueue (0.00s)
    ArrayQueue_test.go:21: &{[6 5 4 3 2] 5 0 5}
    ArrayQueue_test.go:24: 取出两个元素之后
    ArrayQueue_test.go:25: &{[6 5 4 3 2] 5 2 5}
    ArrayQueue_test.go:29: 全部取出之后
    ArrayQueue_test.go:30:
PASS
ok      exercise/ds_algo/study/ArrayQueue       0.282s
*/
