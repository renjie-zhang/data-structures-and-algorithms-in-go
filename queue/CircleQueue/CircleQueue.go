/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-21 19:22:30
 * @LastEditTime: 2019-08-21 19:41:39
 */
package CircleQueue

import "fmt"

type CircleQueue struct {
	q        []interface{}
	capacity int
	rear     int
	front    int
}

//实例化函数
func NewCircleQueue(n int) *CircleQueue {
	if n == 0 {
		return nil
	}
	return &CircleQueue{make([]interface{}, n), n, 0, 0}
}

//队列为空的条件 rear = front
func (this *CircleQueue) IsEmpty() bool {
	if this.rear == this.front {
		return true
	}
	return false
}

//队列已满 (rear+1)%capacity == front
func (this *CircleQueue) IsFull() bool {
	if ((this.rear + 1) % this.capacity) == this.front {
		return true
	}
	return false
}

//数据入队列
func (this *CircleQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.rear] = v
	this.rear = (this.rear + 1) % this.capacity
	return true
}

//数据出队列
func (this *CircleQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return false
	}
	v := this.q[this.front]
	this.front = (this.front + 1) % this.capacity
	return v
}

//展示队列
func (this *CircleQueue) Show() string {
	if this.IsEmpty() {
		return "queue is empty(⊙﹏⊙)"
	}
	result := "front"
	var i = this.front
	for true {
		result += fmt.Sprintf("<-%+v", this.q[i])
		i = (i + 1) % this.capacity
		if i == this.rear {
			break
		}
	}
	result += "<-rear"
	return result
}
