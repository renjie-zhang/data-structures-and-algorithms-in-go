/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-21 19:41:24
 * @LastEditTime: 2019-08-21 19:43:53
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

/*
=== RUN   TestCircleQueue
--- PASS: TestCircleQueue (0.00s)
    CircleQueue_test.go:20: &{[1 2 3 4 <nil>] 5 4 0}
PASS
ok      exercise/ds_algo/study/CircleQueue      0.293s
*/
