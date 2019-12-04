/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-21 19:41:24
 * @LastEditTime: 2019-08-22 09:50:22
 */
package CircleQueue

import "testing"

func TestCircleQueue(t *testing.T) {
	q := NewCircleQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
}
func TestCircleQueue_DeQueue(t *testing.T) {
	q := NewCircleQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	t.Log(q.DeQueue())
	t.Log(q)
	q.EnQueue(5)
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
}

/*
=== RUN   TestCircleQueue
--- PASS: TestCircleQueue (0.00s)
    CircleQueue_test.go:20: &{[1 2 3 4 <nil>] 5 4 0}
=== RUN   TestCircleQueue_DeQueue
--- PASS: TestCircleQueue_DeQueue (0.00s)
    CircleQueue_test.go:30: &{[1 2 3 4 <nil>] 5 4 0}
    CircleQueue_test.go:31: 1
    CircleQueue_test.go:32: &{[1 2 3 4 <nil>] 5 4 1}
    CircleQueue_test.go:34: &{[1 2 3 4 5] 5 0 1}
    CircleQueue_test.go:36: &{[1 2 3 4 5] 5 0 2}
    CircleQueue_test.go:38: &{[1 2 3 4 5] 5 0 3}
    CircleQueue_test.go:40: &{[1 2 3 4 5] 5 0 4}
    CircleQueue_test.go:42: &{[1 2 3 4 5] 5 0 0}
PASS
*/
