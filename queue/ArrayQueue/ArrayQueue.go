package arrayqueue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	front    int
	rear     int
}
//NewArrayQueue
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

//判断是队列是否已满
func (queue *ArrayQueue) IsFull() bool {
	if queue.rear == queue.capacity {
		return true
	}
	return false
}

//判断队列是否为空
func (queue *ArrayQueue) IsEmpty() bool {
	if queue.rear == queue.front {
		return true
	}
	return false
}

//出队列操作
func (queue *ArrayQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	v := queue.q[queue.front]
	queue.front++
	return v
}

//入队操作
func (queue *ArrayQueue) EnQueue(v interface{}) bool {
	if queue.IsFull() {
		return false
	}
	queue.q[queue.rear] = v
	queue.rear++
	return true
}

//打印所有元素
func (queue *ArrayQueue) Show() string {
	if queue.IsEmpty() {
		return "queue is empty(⊙﹏⊙)"
	}
	result := "front"
	for i := queue.front; i <= queue.rear-1; i++ {
		result += fmt.Sprintf("<-%+v", queue.q[i])
	}
	result += "<-rear"
	return result
}
