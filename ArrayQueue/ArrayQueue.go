/*
 * @Descripttion: 使用切片实现队列
 * @version: v1
 * @Author: renjie.zhang
 * @Date: 2019-08-21 18:48:19
 * @LastEditTime: 2019-08-21 19:10:42
 */
package ArrayQueue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	front    int
	rear     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

//判断是队列是否已满
func (this *ArrayQueue) IsFull() bool {
	if this.rear == this.capacity {
		return true
	}
	return false
}

//判断队列是否为空
func (this *ArrayQueue) IsEmpty() bool {
	if this.rear == this.front {
		return true
	}
	return false
}

//出队列操作
func (this *ArrayQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.q[this.front]
	this.front++
	return v
}

//入队操作
func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.rear] = v
	this.rear++
	return true
}

//打印所有元素

func (this *ArrayQueue) Show() string {
	if this.IsEmpty() {
		return "queue is empty(⊙﹏⊙)"
	}
	result := "front"
	for i := this.front; i <= this.rear-1; i++ {
		result += fmt.Sprintf("<-%+v", this.q[i])
	}
	result += "<-rear"
	return result
}
