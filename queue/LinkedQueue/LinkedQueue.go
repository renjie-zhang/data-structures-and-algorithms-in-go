/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-22 18:17:58
 * @LastEditTime: 2019-08-22 18:31:20
 */
package LinkedQueue

type Node struct {
	data int
	next *Node
}

type Queue struct {
	rear *Node
}

//入队操作
func (this *Queue) EnQueue(i int) {
	data := &Node{data: i}
	if this.rear != nil {
		data.next = this.rear
	}
	this.rear = data
}

//出队操作
func (this *Queue) DeQueue() (int, bool) {
	if this.rear == nil {
		return 0, false
	}
	if this.rear.next == nil {
		i := this.rear.data
		this.rear = nil
		return i, true
	}
	current := this.rear
	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, true
		}
		current = current.next
	}
}

func (this *Queue) Peek() (int, bool) {
	if this.rear == nil {
		return 0, false
	}
	return this.rear.data, true
}

func (this *Queue) Get() []int {
	var items []int
	current := this.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}
func (this *Queue) IsEmpty() bool {
	return this.rear == nil
}

func (this *Queue) Empty() {
	this.rear = nil
}
