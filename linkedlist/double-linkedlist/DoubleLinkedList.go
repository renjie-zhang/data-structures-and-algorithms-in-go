package DoubleLinkedList

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

//在链表得头部插入
func (this *LinkedList) InsertFirst(i int) {
	data := &Node{data: i}
	if this.head != nil {
		this.head.prev = data
		data.next = this.head
	}
	this.head = data
}

//在链表得最后插入
func (this *LinkedList) InsertLast(i int) {
	data := &Node{data: i}
	if this.head == nil {
		this.head = data
		this.tail = data
		return
	}
	if this.tail != nil {
		this.tail.next = data
		data.prev = this.tail
	}
	this.tail = data
}

//通过值进行移除
func (this *LinkedList) RemoveByValue(i int) bool {
	if this.head == nil {
		return false
	}
	if this.head.data == i {
		this.head = this.head.next
		this.head.prev = nil
		return true
	}
	if this.tail.data == i {
		this.tail = this.tail.prev
		this.tail.next = nil
		return true
	}
	current := this.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false

}

//通过index进行移除
func (this *LinkedList) RemoveByIndex(index int) bool {
	if this.head == nil {
		return false
	}
	if index < 0 {
		return false
	}
	if index == 0 {
		this.head.prev = nil
		this.head = this.head.next
		return true
	}
	current := this.head
	for j := 1; j < index; j++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

//获得节点个数
func (this *LinkedList) GetSize() int {
	var count int
	current := this.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}
